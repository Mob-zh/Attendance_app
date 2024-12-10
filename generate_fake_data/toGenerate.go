package main

import (
	"attendance_uniapp/initializer"
	"attendance_uniapp/models"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"strconv"
)

func generateFakePassword(plainPassword string) (string, error) {
	// 使用 bcrypt 默认成本生成加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
func generateFakeStudentId() ([]string, error) {
	var ids []string
	minId := 2001
	maxId := 2065
	// 从202200202001到202200202055生成字符串
	for i := minId; i <= maxId; i++ {
		id := "20220020" + strconv.Itoa(i) // 拼接字符串
		ids = append(ids, id)              // 将生成的ID加入切片
	}
	return ids, nil
}
func generateFakeTeacherId() ([]string, error) {
	var ids []string

	// 生成形如 “SZTUBD” + （001~065） 的TeacherId
	for i := 1; i <= 65; i++ {
		id := "SZTUBD" + fmt.Sprintf("%03d", i)
		ids = append(ids, id) // 将生成的ID加入切片
	}
	return ids, nil
}
func main() {
	initializer.Init()
	// 生成假id
	StudentFakeIds, _ := generateFakeStudentId()
	TeacherFakeIds, _ := generateFakeTeacherId()
	// 假设教师和学生各自有一些原始密码
	studentFakePwd := []string{"student123", "student321"}
	teacherFakePwd := []string{"teacher123", "teacher321"}
	var studentFakeHashedPwd []string
	var teacherFakeHashedPwd []string
	// 遍历原始密码列表，生成 bcrypt 加密后的密码
	for _, plainPassword := range studentFakePwd {
		hashedPassword, _ := generateFakePassword(plainPassword)
		studentFakeHashedPwd = append(studentFakeHashedPwd, hashedPassword)
	}
	for _, plainPassword := range teacherFakePwd {
		hashedPassword, _ := generateFakePassword(plainPassword)
		teacherFakeHashedPwd = append(teacherFakeHashedPwd, hashedPassword)
	}
	// 将假数据存入数据库
	for index := range StudentFakeIds {
		//随机二选一密码
		randomIndex := rand.Intn(len(studentFakeHashedPwd))
		initializer.DB.Create(&models.Student{StudentId: StudentFakeIds[index], StudentName: faker.ChineseName(), StudentPwd: studentFakeHashedPwd[randomIndex]})
	}
	for _, id := range TeacherFakeIds {
		//随机二选一密码
		randomIndex := rand.Intn(len(teacherFakeHashedPwd))
		initializer.DB.Create(&models.Teacher{TeacherId: id, TeacherName: faker.ChineseName(), TeacherPwd: teacherFakeHashedPwd[randomIndex]})
	}

	log.Println("Generate complete!")
}
