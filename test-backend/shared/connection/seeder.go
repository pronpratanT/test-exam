package connection

import (
	"log"
	"test-backend/shared/models"

	"gorm.io/gorm"
)

func SeedMockDataTest3(db *gorm.DB) {
	// 1. ตรวจสอบให้แน่ใจว่าได้สร้างตารางใน Database ตามโครงสร้างของ models.Test3 แล้ว
	err := db.AutoMigrate(&models.Test3{})
	if err != nil {
		log.Println("Failed to migrate Test3:", err)
		return
	}

	// 2. ตรวจสอบว่ามีข้อมูลอยู่ในตาราง Test3 หรือยัง
	var count int64
	db.Model(&models.Test3{}).Count(&count)

	if count == 0 {
		// 3. ถ้าไม่มีข้อมูลเลย ให้เตรียม Mockup Data เป็น Array
		mockData := []models.Test3{
			{Name: "Somchai Jaidee", Reason: "เบิกค่าเดินทาง", Status: "pending"},
			{Name: "Suda Yindee", Reason: "ลาป่วย", Status: "pending"},
			{Name: "Mana Rakthai", Reason: "อุปกรณ์ IT เสียหาย", Status: "pending"},
			{Name: "Anan Dee", Reason: "ลากิจ", Status: "pending"},
			{Name: "Kanya Sukjai", Reason: "ค่าอาหารลูกค้า", Status: "pending"},
			{Name: "Prasert Meechai", Reason: "ลาพักร้อน", Status: "pending"},
		}

		// 4. สั่ง Insert ทีเดียวทั้งหมดเข้าไปที่ Database
		if err := db.Create(&mockData).Error; err != nil {
			log.Println("Failed to seed Test3 mock data:", err)
		} else {
			log.Println("Successfully seeded Test3 mock data!")
		}
	} else {
		log.Println("Test3 table already has data. Skipping seed.")
	}
}
