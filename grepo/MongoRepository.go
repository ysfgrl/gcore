package grepo

import (
	"context"
	"slices"
	"strings"
	"time"

	"github.com/ysfgrl/gcore/gerror"
	"github.com/ysfgrl/gcore/gmodel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var createdAtKey = "createdAt"
var idKey = "_id"
var dnaKey = "dna"

var constKeys = []string{
	createdAtKey,
	idKey,
	dnaKey,
}

type Repository[DType any] struct {
	Collection    *mongo.Collection
	FilterKeys    []string
	AggregatePipe []bson.M
	Dna           string
}

func (repo *Repository[DType]) GetById(ctx context.Context, id primitive.ObjectID, isAgg bool) (*DType, *gerror.Error) {

	return repo.GetByQuery(ctx, bson.M{
		idKey:  id,
		dnaKey: primitive.Regex{Pattern: "^" + repo.Dna, Options: "i"},
	}, isAgg)

}

func (repo *Repository[DType]) GetByKey(ctx context.Context, key string, value any, isAgg bool) (*DType, *gerror.Error) {
	return repo.GetByQuery(ctx, bson.M{
		key: value,
	}, isAgg)
}

func (repo *Repository[DType]) GetByQuery(ctx context.Context, query bson.M, isAgg bool) (*DType, *gerror.Error) {
	if len(repo.Dna) > 0 {
		query[dnaKey] = repo.Dna
	}
	var item DType
	if !isAgg {
		res := repo.Collection.FindOne(ctx, query)
		if err := res.Decode(&item); err != nil {
			return nil, gerror.GetError(err)
		}
		return &item, nil
	}
	return repo.GetByAggregate(ctx, query, repo.AggregatePipe)
}

func (repo *Repository[DType]) GetByAggregate(ctx context.Context, query bson.M, agg []bson.M) (*DType, *gerror.Error) {
	var item DType
	pipeline := append(agg, bson.M{"$match": query})
	pipeline = append(pipeline, bson.M{"$limit": 1})
	pipeline = append(pipeline, bson.M{"$skip": 0})
	cur, err := repo.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, gerror.GetError(err)
	}
	defer cur.Close(ctx)
	if cur.Next(ctx) {
		err = cur.Decode(&item)
		if err != nil {
			return nil, gerror.GetError(err)
		}
		return &item, nil
	}
	return nil, &gerror.Error{
		Code:   "mongo.notfound",
		Detail: "Document not found",
	}
}

func (repo *Repository[DType]) List(ctx context.Context, filters gmodel.ListRequest) (*gmodel.ListResponse[DType], *gerror.Error) {
	query := bson.M{}
	gte, gteOk := filters.Filters["gte"]
	lte, lteOk := filters.Filters["lte"]
	if gteOk && lteOk {
		query[createdAtKey] = bson.M{"$gte": gte, "$lt": lte}
	} else if lteOk {
		query[createdAtKey] = bson.M{"$lt": lte}
	} else if gteOk {
		query[createdAtKey] = bson.M{"$gte": gte}
	}

	for key, val := range filters.Filters {
		if key == "gte" || key == "lte" {
			continue
		}
		if key == "keyword" {
			if len(repo.FilterKeys) == 0 {
				continue
			}
			and := bson.A{}
			for _, filterKey := range repo.FilterKeys {
				and = append(and, bson.M{filterKey: primitive.Regex{Pattern: val.(string), Options: "i"}})
			}
			query["$or"] = and
		} else if key == "dna" {
			query[key] = primitive.Regex{Pattern: "^" + val.(string), Options: "i"}
		} else if strings.HasSuffix(key, "Id") {
			switch val.(type) {
			case string:
				id, err := primitive.ObjectIDFromHex(val.(string))
				if err == nil {
					query[key] = id
				}
			case primitive.ObjectID:
				query[key] = val.(primitive.ObjectID)
			}
		} else {
			query[key] = val
		}
	}
	pipeline := append(repo.AggregatePipe, bson.M{"$match": query})
	pipeline = append(pipeline, bson.M{"$sort": bson.M{
		createdAtKey: -1,
	}})
	pipeline = append(pipeline, bson.M{"$skip": int64((filters.Page - 1) * filters.PageSize)})
	pipeline = append(pipeline, bson.M{"$limit": filters.PageSize})

	cur, err := repo.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, gerror.GetError(err)
	}
	defer cur.Close(ctx)
	var list []DType
	for cur.Next(ctx) {
		var item DType
		err := cur.Decode(&item)
		if err != nil {
			return nil, gerror.GetError(err)
		}
		list = append(list, item)
	}
	if err := cur.Err(); err != nil {
		return nil, gerror.GetError(err)
	}
	count, err := repo.Collection.CountDocuments(ctx, query)
	if err != nil {
		return nil, gerror.GetError(err)
	}
	if len(list) == 0 {
		list = []DType{}
	}
	return &gmodel.ListResponse[DType]{
		Page:     filters.Page,
		PageSize: filters.PageSize,
		Total:    count,
		List:     list,
	}, nil
}

func (repo *Repository[DType]) Create(ctx context.Context, schema DType) (*DType, *gerror.Error) {

	value, err := bson.Marshal(schema)
	if err != nil {
		return nil, gerror.GetError(err)
	}
	query := bson.M{}
	if err := bson.Unmarshal(value, query); err != nil {
		return nil, gerror.GetError(err)
	}
	query[createdAtKey] = time.Now().UTC()
	res, err := repo.Collection.InsertOne(ctx, query)
	if err != nil {
		return nil, gerror.GetError(err)
	}
	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, gerror.UserError("create.error", "error")
	}
	return repo.GetById(ctx, id, false)
}

func (repo *Repository[DType]) CreateMany(ctx context.Context, schemas []DType) (int, *gerror.Error) {

	var queries = make([]interface{}, len(schemas))
	now := time.Now().UTC()
	for i := range schemas {
		value, err := bson.Marshal(schemas[i])
		if err != nil {
			return 0, gerror.GetError(err)
		}
		query := bson.M{}
		if err = bson.Unmarshal(value, query); err != nil {
			return 0, gerror.GetError(err)
		}
		query[createdAtKey] = now
		queries[i] = query
	}
	res, err := repo.Collection.InsertMany(ctx, queries)
	if err != nil {
		return 0, gerror.GetError(err)
	}
	return len(res.InsertedIDs), nil
}
func (repo *Repository[DType]) Increment(ctx context.Context, id primitive.ObjectID, key string, val int) (bool, *gerror.Error) {
	opts := options.Update().SetUpsert(false)
	_, err := repo.Collection.UpdateOne(
		ctx,
		bson.D{{"_id", id}},
		bson.D{{"$inc", bson.M{
			key: val,
		}}},
		opts)
	if err != nil {
		return false, gerror.GetError(err)
	}
	return true, nil
}

func (repo *Repository[DType]) UpdateById(ctx context.Context, id primitive.ObjectID, schema DType) *gerror.Error {

	raw, err := bson.Marshal(schema)
	if err != nil {
		return gerror.GetError(err)
	}
	update := bson.M{}
	if err := bson.Unmarshal(raw, &update); err != nil {
		return gerror.GetError(err)
	}
	for i := range constKeys {
		delete(update, constKeys[i])
	}

	_, err = repo.Collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.D{{"$set", update}},
	)
	if err != nil {
		return gerror.GetError(err)
	}
	return nil
}

func (repo *Repository[DType]) UpdateField(ctx context.Context, id primitive.ObjectID, field string, value any) *gerror.Error {

	if slices.Contains(constKeys, field) {
		return nil
	}
	_, err := repo.Collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.D{{"$set", bson.D{{field, value}}}},
	)
	if err != nil {
		return gerror.GetError(err)
	}
	return nil
}
func (repo *Repository[DType]) UpdateFields(ctx context.Context, id primitive.ObjectID, fields map[string]any) *gerror.Error {

	for key := range fields {
		if slices.Contains(constKeys, key) {
			delete(fields, key)
		}
	}
	if len(fields) == 0 {
		return nil
	}
	set := bson.D{}
	for key, value := range fields {
		set = append(set, bson.E{Key: key, Value: value})
	}
	_, err := repo.Collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.D{{"$set", set}},
	)
	if err != nil {
		return gerror.GetError(err)
	}
	return nil
}

func (repo *Repository[DType]) DeleteById(ctx context.Context, id primitive.ObjectID) *gerror.Error {
	_, err := repo.Collection.DeleteOne(
		ctx,
		bson.M{idKey: id},
	)
	if err != nil {
		return gerror.GetError(err)
	}
	return nil
}

func (repo *Repository[DType]) Count(ctx context.Context, query bson.D) (int64, *gerror.Error) {
	opts := options.Count().SetHint("_id_")
	count, err := repo.Collection.CountDocuments(ctx, query, opts)
	if err != nil {
		return 0, gerror.GetError(err)
	}
	return count, nil
}
