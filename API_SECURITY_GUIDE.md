# 🔐 คู่มือการใช้งานระบบความปลอดภัย API

## ภาพรวม

ระบบ API ได้รับการป้องกันด้วยระบบ Authentication และ Authorization เพื่อป้องกันไม่ให้บุคคลทั่วไปเข้าถึง API ได้โดยไม่ได้รับอนุญาต

## วิธีการทำงาน

### 1. การเข้าสู่ระบบ (Login)
```bash
POST /authEntry/login
Content-Type: application/json

{
    "username": "admin",
    "password": "123456"
}
```

**Response:**
```json
{
    "success": true,
    "message": "✅ เข้าสู่ระบบสำเร็จ",
    "role": "admin",
    "token": "eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImFkbWluIiwicm9sZSI6ImFkbWluIiwiZXhwIjoxNzM1NjgwMDAwfQ=="
}
```

### 2. การใช้ Token

หลังจากได้รับ token แล้ว ให้ส่ง token ไปกับทุก request ในรูปแบบ:

```bash
Authorization: Bearer eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImFkbWluIiwicm9sZSI6ImFkbWluIiwiZXhwIjoxNzM1NjgwMDAwfQ==
```

## ระดับการเข้าถึง

### 🔓 Public Routes (ไม่ต้อง Authentication)
- `POST /authEntry/login` - เข้าสู่ระบบ
- `POST /authEntry/registerUser` - ลงทะเบียนผู้ใช้ทั่วไป
- `POST /authEntry/registerAdmin` - ลงทะเบียนผู้ดูแลระบบ
- `POST /authEntry/logout` - ออกจากระบบ
- `GET /healthEntry/health` - ตรวจสอบสถานะระบบ

### 🔐 Protected Routes (ต้อง Authentication)
- `POST /problemEntry/reportProblem` - รายงานปัญหา
- `PUT /problemEntry/solveProblem` - แก้ไขปัญหา
- `GET /problemEntry/problems` - ดึงรายการปัญหา
- `GET /problemEntry/problem/{id}` - ดึงปัญหาตาม ID
- `PUT /problemEntry/problem/{id}` - อัปเดตปัญหา
- `DELETE /problemEntry/problem/{id}` - ลบปัญหา
- `PUT /problemEntry/problem/{id}/reset-solution` - รีเซ็ตการแก้ไข
- `PUT /problemEntry/problem/{id}/update-solution` - อัปเดตการแก้ไข
- `PUT /problemEntry/problem/{id}/update-problem` - อัปเดตปัญหาและการแก้ไข
- `DELETE /problemEntry/problem/{id}/delete-solution` - ลบการแก้ไข
- `POST /branchEntry/branchOffice` - เพิ่มสาขา
- `PUT /branchEntry/branchOffice/{ip_phone}` - อัปเดตสาขา
- `DELETE /branchEntry/branchOffice/{ip_phone}` - ลบสาขา
- `GET /branchEntry/branchOffices` - ดึงรายการสาขา
- `POST /programEntry/program` - เพิ่มโปรแกรม
- `GET /programEntry/programs` - ดึงรายการโปรแกรม
- `PUT /programEntry/program/{id}` - อัปเดตโปรแกรม
- `DELETE /programEntry/program/{id}` - ลบโปรแกรม
- `GET /dashboardEntry/dashboard` - ดึงข้อมูลแดชบอร์ด

### 👑 Admin Only Routes (ต้องเป็น Admin)
- `POST /userEntry/user` - เพิ่มผู้ใช้
- `GET /userEntry/users` - ดึงรายการผู้ใช้
- `PUT /authEntry/updateUser` - อัปเดตผู้ใช้
- `DELETE /authEntry/deleteUser` - ลบผู้ใช้
- `DELETE /problemEntry/deleteAllProblems` - ลบปัญหาทั้งหมด
- `DELETE /branchEntry/deleteAllBranchOffices` - ลบสาขาทั้งหมด
- `DELETE /programEntry/deleteAllPrograms` - ลบโปรแกรมทั้งหมด

## ตัวอย่างการใช้งาน

### 1. เข้าสู่ระบบ
```bash
curl -X POST http://localhost:5000/authEntry/login \
  -H "Content-Type: application/json" \
  -d '{"username": "admin", "password": "123456"}'
```

### 2. ใช้ Token เพื่อเข้าถึง API
```bash
curl -X GET http://localhost:5000/problemEntry/problems \
  -H "Authorization: Bearer eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImFkbWluIiwicm9sZSI6ImFkbWluIiwiZXhwIjoxNzM1NjgwMDAwfQ=="
```

### 3. เข้าถึง Admin API
```bash
curl -X GET http://localhost:5000/userEntry/users \
  -H "Authorization: Bearer eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImFkbWluIiwicm9sZSI6ImFkbWluIiwiZXhwIjoxNzM1NjgwMDAwfQ=="
```

## ข้อผิดพลาดที่อาจเกิดขึ้น

### 401 Unauthorized
- ไม่มี Authorization header
- Token ไม่ถูกต้อง
- Token หมดอายุ

### 403 Forbidden
- ไม่มีสิทธิ์เข้าถึง (ต้องเป็น Admin)

## การตั้งค่า

### Environment Variables
สร้างไฟล์ `.env` ในโฟลเดอร์หลัก:

```env
# Database Configuration
DB_USER=root
DB_PASS=123456
DB_HOST=192.168.1.153
DB_PORT=3306
DB_NAME=MySQLdatabases

# JWT Configuration
JWT_SECRET=your-super-secret-jwt-key-change-this-in-production

# Server Configuration
PORT=5000

# CORS Configuration
CORS_ALLOWED_ORIGINS=*
```

## ความปลอดภัย

1. **Token Expiration**: Token มีอายุ 24 ชั่วโมง
2. **Role-based Access**: แยกสิทธิ์ระหว่าง User และ Admin
3. **Secure Headers**: ใช้ Authorization header แบบ Bearer token
4. **Input Validation**: ตรวจสอบข้อมูลที่รับเข้ามา

## การทดสอบ

### ทดสอบการเข้าถึงโดยไม่มี Token
```bash
curl -X GET http://localhost:5000/problemEntry/problems
# ควรได้ 401 Unauthorized
```

### ทดสอบการเข้าถึง Admin API ด้วย User Token
```bash
# ใช้ token ของ user ที่ไม่ใช่ admin
curl -X GET http://localhost:5000/userEntry/users \
  -H "Authorization: Bearer [user-token]"
# ควรได้ 403 Forbidden
```

## การแก้ไขปัญหา

### Token หมดอายุ
หากได้รับข้อผิดพลาด "Token expired" ให้เข้าสู่ระบบใหม่เพื่อรับ token ใหม่

### ไม่มีสิทธิ์เข้าถึง
หากต้องการเข้าถึง Admin API ให้ใช้บัญชีที่มี role เป็น "admin" 