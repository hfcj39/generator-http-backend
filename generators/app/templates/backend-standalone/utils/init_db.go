package utils

import "<%= displayName %>/model"

func InitDB(functions ...model.InitDBFunc) (err error) {
	for _, v := range functions {
		err = v.Init()
		if err != nil {
			return err
		}
	}
	return nil
}
