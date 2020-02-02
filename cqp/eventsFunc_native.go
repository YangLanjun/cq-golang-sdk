// Code generated by cqgen, DO NOT EDIT.
// +build !websocket

package cqp

// #include <stdint.h>
import "C"

//export _appinfo
func _appinfo() *C.char { return C.CString("9," + AppID) }

//export _on_enable
func _on_enable() int32 {
	if Enable == nil {
		return 0
	}
	defer panicToFatal()
	return Enable()
}

//export _on_disable
func _on_disable() int32 {
	if Disable == nil {
		return 0
	}
	defer panicToFatal()
	return Disable()
}

//export _on_start
func _on_start() int32 {
	if Start == nil {
		return 0
	}
	defer panicToFatal()
	return Start()
}

//export _on_exit
func _on_exit() int32 {
	if Exit == nil {
		return 0
	}
	defer panicToFatal()
	return Exit()
}

//export _on_private_msg
func _on_private_msg(var0, var1 C.int32_t, var2 C.int64_t, var3 *C.char, var4 C.int32_t) int32 {
	if PrivateMsg == nil {
		return 0
	}
	defer panicToFatal()
	return PrivateMsg(int32(var0), int32(var1), int64(var2), goString(var3), int32(var4))
}

//export _on_group_msg
func _on_group_msg(var0, var1 C.int32_t, var2, var3 C.int64_t, var4, var5 *C.char, var6 C.int32_t) int32 {
	if GroupMsg == nil {
		return 0
	}
	defer panicToFatal()
	return GroupMsg(int32(var0), int32(var1), int64(var2), int64(var3), goString(var4), goString(var5), int32(var6))
}

//export _on_discuss_msg
func _on_discuss_msg(var0, var1 C.int32_t, var2, var3 C.int64_t, var4 *C.char, var5 C.int32_t) int32 {
	if DiscussMsg == nil {
		return 0
	}
	defer panicToFatal()
	return DiscussMsg(int32(var0), int32(var1), int64(var2), int64(var3), goString(var4), int32(var5))
}

//export _on_group_upload
func _on_group_upload(var0, var1 C.int32_t, var2, var3 C.int64_t, var4 *C.char) int32 {
	if GroupUpload == nil {
		return 0
	}
	defer panicToFatal()
	return GroupUpload(int32(var0), int32(var1), int64(var2), int64(var3), goString(var4))
}

//export _on_group_admin
func _on_group_admin(var0, var1 C.int32_t, var2, var3 C.int64_t) int32 {
	if GroupAdmin == nil {
		return 0
	}
	defer panicToFatal()
	return GroupAdmin(int32(var0), int32(var1), int64(var2), int64(var3))
}

//export _on_group_member_decrease
func _on_group_member_decrease(var0, var1 C.int32_t, var2, var3, var4 C.int64_t) int32 {
	if GroupMemberDecrease == nil {
		return 0
	}
	defer panicToFatal()
	return GroupMemberDecrease(int32(var0), int32(var1), int64(var2), int64(var3), int64(var4))
}

//export _on_group_member_increase
func _on_group_member_increase(var0, var1 C.int32_t, var2, var3, var4 C.int64_t) int32 {
	if GroupMemberIncrease == nil {
		return 0
	}
	defer panicToFatal()
	return GroupMemberIncrease(int32(var0), int32(var1), int64(var2), int64(var3), int64(var4))
}

//export _on_friend_add
func _on_friend_add(var0, var1 C.int32_t, var2 C.int64_t) int32 {
	if FriendAdd == nil {
		return 0
	}
	defer panicToFatal()
	return FriendAdd(int32(var0), int32(var1), int64(var2))
}

//export _on_friend_request
func _on_friend_request(var0, var1 C.int32_t, var2 C.int64_t, var3, var4 *C.char) int32 {
	if FriendRequest == nil {
		return 0
	}
	defer panicToFatal()
	return FriendRequest(int32(var0), int32(var1), int64(var2), goString(var3), goString(var4))
}

//export _on_group_request
func _on_group_request(var0, var1 C.int32_t, var2, var3 C.int64_t, var4, var5 *C.char) int32 {
	if GroupRequest == nil {
		return 0
	}
	defer panicToFatal()
	return GroupRequest(int32(var0), int32(var1), int64(var2), int64(var3), goString(var4), goString(var5))
}
