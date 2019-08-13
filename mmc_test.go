package main

import "testing"

func TestSelect(t *testing.T) {
	args := []string{"192.168.1.10", "root", "root", "select * apple.user book.user;"}
	DBTool(args)
}
func TestSelect2(t *testing.T) {
	args := []string{"192.168.1.10", "root", "root", "select * from mysql.user;"}
	DBTool(args)
}
func TestInsert(t *testing.T) {
	args := []string{"192.168.1.10", "root", "root", "insert into apple.user(id,name,age,account) values(3,'Joe',24,99.99);"}
	DBTool(args)
}

func TestCreateDatabase(t *testing.T) {
	args := []string{"192.168.1.10", "root", "root", "CREATE DATABASE IF NOT EXISTS apple DEFAULT CHARSET utf8mb4 COLLATE = utf8mb4_unicode_ci;"}
	DBTool(args)
}
func TestShowDatabase(t *testing.T) {
	args := []string{"192.168.1.10", "root", "root", "show databases;"}
	DBTool(args)
}
