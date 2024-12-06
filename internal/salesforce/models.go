package salesforce

const (
	OperandEqual            = "equal"
	OperandNotEqual         = "notEqual"
	OperandLessThan         = "lessThan"
	OperandLessThanEqual    = "lessThanOrEqual"
	OperandGreaterThan      = "greaterThan"
	OperandGreaterThanEqual = "greaterThanOrEqual"
	OperandLike             = "like"
	OperandIsNull           = "isNull"
	OperandIsNotNull        = "isNotNull"
	OperandContains         = "contains"
	OperandMustContain      = "mustContain"
	OperandStartsWith       = "startsWith"
	OperandIn               = "in"
)

const (
	LogicalOperandAnd = "AND"
	LogicalOperandOr  = "OR"
)

const (
	SortDirectionAscending  = "ASC"
	SortDirectionDescending = "DESC"
)

type PageObject struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"pageSize,omitempty"`
}

type QueryOperand struct {
	Property       string `json:"property,omitempty"`
	SimpleOperator string `json:"simpleOperator,omitempty"`
	Value          any    `json:"value,omitempty"`
}

type QueryObject struct {
	LeftOperand     QueryOperand `json:"leftOperand,omitempty"`
	RightOperand    QueryOperand `json:"rightOperand,omitempty"`
	LogicalOperator string       `json:"logicalOperator"`
}

type SortObject struct {
	Property  string `json:"propery,omitempty"`
	Direction string `json:"direction,omitempty"`
}
