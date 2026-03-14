package contact

// Department 部门信息
type Department struct {
	// ID 部门id
	ID int `json:"id"`
	// Name 部门名称
	Name string `json:"name"`
	// NameEN 部门英文名称
	NameEN string `json:"name_en,omitempty"`
	// DepartmentLeader 部门负责人的UserID列表
	DepartmentLeader []string `json:"department_leader,omitempty"`
	// ParentID 父部门id
	ParentID int `json:"parentid"`
	// Order 在父部门中的次序�?
	Order int `json:"order,omitempty"`
}

