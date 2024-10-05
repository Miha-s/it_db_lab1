package attributes

func CreateAttribute(attr_type string, attr_name string) Attribute {
	switch attr_type {
	case "char":
		return NewCharAttribute(attr_name)
	case "int":
		return NewIntegerAttribute(attr_name)
	case "real":
		return NewRealAttribute(attr_name)
	default:
		panic("unknown data type: " + attr_type)
	}
}
