/**
 *******************************************slade********************************************
 * Copyright (c)  slade
 * Created by super-class.
 * User: 605724193@qq.com
 * Date: 2019/08/07
 * Time: 11:18
 ********************************************************************************************
 */

package user

import (
	"errors"

	"super-class/util"
)

var (
	ErrUserNotExist         = util.AppErrorNew(errors.New("user not exist"), util.ERROR)
)
