# gcore

A modular and extensible Go core library providing essential building blocks for scalable backend services.  
This project includes modules for error handling, models, repository pattern, storage integrations, routing, authentication, and configuration management.

---

## 📦 Modules Overview

### **gerror** – Error Handling
Provides a structured error system with codes, messages, and stack-friendly helpers for consistent API error responses.

### **gmodel** – Base Models
Contains shared base structs for request/response models to ensure consistent formatting across services.

### **grepo** – MongoDB Repository Layer
Implements a clean repository pattern for MongoDB, including:
- BaseRepo with CRUD helpers
- Context-aware operations
- Flexible query utilities

### **gstorage** – MinIO Integration
Provides a unified storage interface with:
- Upload / Download utilities
- Presigned URL support
- Bucket management

### **groute** – Fiber Routing Controller
Opinionated controllers for the Fiber framework:
- Route grouping
- Auto-binding
- Middleware support

### **gauth** – JWT Authentication
Features include:
- Token generation & validation
- Claims management
- Middleware helpers for Fiber

### **gconf** – Configuration Loader
Reads environment variables/config files and maps them into structured Go types.

---

## 🚀 Installation

```sh
go get github.com/ysfgrl/gcore
```

To import individual modules:

```go
import "github.com/ysfgrl/gcore/gerror"
import "github.com/ysfgrl/gcore/gmodel"
import "github.com/ysfgrl/gcore/grepo"
import "github.com/ysfgrl/gcore/gstorage"
import "github.com/ysfgrl/gcore/groute"
import "github.com/ysfgrl/gcore/gauth"
import "github.com/ysfgrl/gcore/gconf"
```

---

## 🧩 Basic Usage Examples

### **gerror Example**

```go


```

### **gmodel Example**

```go


```

### **grepo Example**

```go


```

### **gstorage Example**

```go


```

### **groute Example**

```go


```

### **gauth Example**

```go

import "github.com/gofiber/fiber/v2"
import "github.com/golang-jwt/jwt/v5"
import "github.com/ysfgrl/gcore/gauth"
import "github.com/ysfgrl/gcore/gerror"

var privateKey *rsa.PrivateKey = nil
var publicKey *rsa.PublicKey = ni

var IJWTAuth gauth.IAuth = nil

type JWTAuth struct {
    gauth.BaseAuth
}

func (a *JWTAuth) Init() {
//init 
}

func init() {
    IJWTAuth = &JWTAuth{
        BaseAuth: gauth.BaseAuth{
        PrivateKey: privateKey,
        PublicKey:  publicKey,
        AuthScheme: "Bearer",
        Method:      jwt.SigningMethodRS256,
        TokenLookup: "header:Authorization",
    },
}
app := fiber.New()
    app.Post("/test", IJWTAuth.Require, Test)
}


func Test(c *fiber.Ctx) error {
    return c.Status(200).JSON(fiber.Map{})
}

```

### **gconf Example**

```go


```

---

## 🗂 Project Structure

```
gcore/
├── gerror/
├── gmodel/
├── grepo/
├── gstorage/
├── groute/
├── gauth/
├── gconf/
└── README.md
```

---

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Submit a pull request

---

## 📄 License

MIT License © 2025 ysfgrl

---

## ⭐ Support

If you find this project helpful, please star the repository!  
https://github.com/ysfgrl/gcore
