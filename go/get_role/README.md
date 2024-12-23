# 回家作業說明

## 實作主題

霹靂布袋戲角色資料庫系統

## 事前準備

- 請使用事先建立好的 Go 專案範本進行開發

## 實作目標

請利用 `gin` 實作五個 RESTful APIs

- 取得全部資料 [GET] http://localhost:8080/role
- 新增單筆資料 [POST] http://localhost:8080/role
- 取得單筆資料 [GET] http://localhost:8080/role/:id
- 更新單筆資料 [PUT] http://localhost:8080/role/:id
- 刪除單筆資料 [DELETE] http://localhost:8080/role/:id

## 實作方式

請直接利用 `slice` 管理資料集合，無須透過資料庫存取！
