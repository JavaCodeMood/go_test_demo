package main

import (
	"fmt"
	"strings"
		)

//截取手机号码的最后一位

func main(){
	var mobile string = "18295514403"

	//截取字符串
	content := mobile[10 : len(mobile)]
	fmt.Println(content)
	fmt.Println("t_order_",content)

	//拼接字符串
	s := "t_order_"
	fmt.Println(strings.Join([]string{s,content},""))

	fmt.Println(strings.Join([]string{"t_order_",content},""))

	fmt.Println()



}


/*if content == string(0) {
		rows, err := rederp.DB.Query(`select tp.product_name, tor.total_fee, tor.quantity,tor.create_time,tor.channel,tor.order_id from stock_app.t_order_0 tor, stock_app.t_products tp WHERE tor.product_id = tp.product_id AND tor.status = 9000 AND tor.userid = ?`, params.UserId)
	}else if content == string(1){
		rows, err := rederp.DB.Query(`select tp.product_name, tor.total_fee, tor.quantity,tor.create_time,tor.channel,tor.order_id from stock_app.t_order_1 tor, stock_app.t_products tp WHERE tor.product_id = tp.product_id AND tor.status = 9000 AND tor.userid = ?`, params.UserId)
	}else if content == string(2){
		rows, err := rederp.DB.Query(`select tp.product_name, tor.total_fee, tor.quantity,tor.create_time,tor.channel,tor.order_id from stock_app.t_order_2 tor, stock_app.t_products tp WHERE tor.product_id = tp.product_id AND tor.status = 9000 AND tor.userid = ?`, params.UserId)
	}else if content == string(3){
		rows, err := rederp.DB.Query(`select tp.product_name, tor.total_fee, tor.quantity,tor.create_time,tor.channel,tor.order_id from stock_app.t_order_3 tor, stock_app.t_products tp WHERE tor.product_id = tp.product_id AND tor.status = 9000 AND tor.userid = ?`, params.UserId)
	}else if content == string(4){
		rows, err := rederp.DB.Query(`select tp.product_name, tor.total_fee, tor.quantity,tor.create_time,tor.channel,tor.order_id from stock_app.t_order_4 tor, stock_app.t_products tp WHERE tor.product_id = tp.product_id AND tor.status = 9000 AND tor.userid = ?`, params.UserId)
	}else if content == string(5){
		rows, err := rederp.DB.Query(`select tp.product_name, tor.total_fee, tor.quantity,tor.create_time,tor.channel,tor.order_id from stock_app.t_order_5 tor, stock_app.t_products tp WHERE tor.product_id = tp.product_id AND tor.status = 9000 AND tor.userid = ?`, params.UserId)
	}else if content == string(6){
		rows, err := rederp.DB.Query(`select tp.product_name, tor.total_fee, tor.quantity,tor.create_time,tor.channel,tor.order_id from stock_app.t_order_6 tor, stock_app.t_products tp WHERE tor.product_id = tp.product_id AND tor.status = 9000 AND tor.userid = ?`, params.UserId)
	}else if content == string(7){
		rows, err := rederp.DB.Query(`select tp.product_name, tor.total_fee, tor.quantity,tor.create_time,tor.channel,tor.order_id from stock_app.t_order_7 tor, stock_app.t_products tp WHERE tor.product_id = tp.product_id AND tor.status = 9000 AND tor.userid = ?`, params.UserId)
	}else if content == string(8){
		rows, err := rederp.DB.Query(`select tp.product_name, tor.total_fee, tor.quantity,tor.create_time,tor.channel,tor.order_id from stock_app.t_order_8 tor, stock_app.t_products tp WHERE tor.product_id = tp.product_id AND tor.status = 9000 AND tor.userid = ?`, params.UserId)
	}else if content == string(9){
		rows, err := rederp.DB.Query(`select tp.product_name, tor.total_fee, tor.quantity,tor.create_time,tor.channel,tor.order_id from stock_app.t_order_9 tor, stock_app.t_products tp WHERE tor.product_id = tp.product_id AND tor.status = 9000 AND tor.userid = ?`, params.UserId)
	}*/

/*
func (h *User) QueryUserInfo(r *http.Request, params *GetUserInfoParam, info *UserRegistInfo) error {
	var buff sql.NullString
	_ = gpxj.DB.QueryRow(`SELECT userid AS mobile, nickname, registtime, invite_code, invite_num FROM stock_app.t_user_base WHERE userid = ?`, params.Userid).Scan(&buff)

	var InviteNum int
	if buff.Valid {
		info.Userid = buff.String
		info.Nickname = buff.String
		info.Registtime = buff.String
		info.Invite_code = buff.String
		info.Invite_num = InviteNum
	}

	return nil
}
*/