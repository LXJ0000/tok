package utils

import "github.com/bwmarrin/snowflake"

var defaultNode *snowflake.Node

func InitDefaultSnowflakeNode(node int64) {
	var err error
	defaultNode, err = snowflake.NewNode(node)
	if err != nil {
		panic(err)
	}
}

func GenID() int64 {
	return defaultNode.Generate().Int64()
}
