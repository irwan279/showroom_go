package utils

const (
	// USER
	INSERT_USER      = "INSERT INTO ms_user(id,username, password, role, is_active ) VALUES($1, $2, $3, $4, $5)"
	GET_USER_BY_ID   = "SELECT id, username,role,is_active FROM ms_user WHERE id = $1"
	GET_ALL_USER     = "SELECT id, username,role,is_active FROM ms_user"
	GET_USER_BY_NAME = "SELECT id, username,password,role,is_active FROM ms_user WHERE username = $1"
	EDIT_USER_ID     = "UPDATE ms_user SET username=$1, password=$2, is_active=$3 WHERE id = $4"

	// Vehicle
	DELETE_VEHICLE        = "DELETE FROM ms_vehicle WHERE id = $1;"
	UPDATE_VEHICLE        = "UPDATE ms_vehicle SET name=$1, type=$2, identification_number=$3, machine_number=$4, release_date=$5, price=$6, price_rent=$7, status=$8, is_available=$9, number_plate=$10, stnk=$11, no_bpkb=$12 WHERE id=$13"
	GET_ALL_VEHICLE       = "SELECT id, name, type, identification_number, machine_number, release_date, price, status, is_available, number_plate, stnk, no_bpkb,  price_rent  FROM ms_vehicle ORDER BY id ASC;"
	GET_ALL_VEHICLE_RENT  = "SELECT id, name, type, identification_number, machine_number, release_date, price, status, is_available, number_plate, stnk, no_bpkb FROM ms_vehicle WHERE is_available = false"
	GET_ALL_VEHICLE_READY = "SELECT id, name, type, identification_number, machine_number, release_date, price, status, is_available, number_plate, stnk, no_bpkb FROM ms_vehicle WHERE is_available = true"
	GET_VEHICLE_BY_NAME   = "SELECT id, name, type, identification_number, machine_number, release_date, price, status, is_available, number_plate, stnk, no_bpkb, price_rent FROM ms_vehicle WHERE name = $1"
	GET_VEHICLE_BY_ID     = "SELECT id, name, type, identification_number, machine_number, release_date, price, status, is_available, number_plate, stnk, no_bpkb, price_rent FROM ms_vehicle WHERE id = $1"
	INSERT_VEHICLE        = "INSERT INTO ms_vehicle (id, name, type, identification_number, machine_number, release_date, price, status, is_available, number_plate, stnk, no_bpkb, price_rent) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13);"

	//cash
	INSERT_CASH_STATEMENT    = "insert into tx_cash(id, id_vehicle, id_customer, price, date_payment, created_at, created_by) values($1,$2,$3,$4,$5,$6,$7)"
	INSERT_CASH_S_UPDATE_VHC = "Update ms_vehicle SET is_available=$1 WHERE id =$2"
	GET_CASH_ID              = "select id, id_vehicle, id_customer, price, date_payment, created_at, created_by from tx_cash where id=$1"
	GET_CASH_ALL             = "select id, id_vehicle, id_customer, price, date_payment, created_at, created_by from tx_cash"

	//customer
	INSERT_CUST        = "INSERT INTO ms_user(id,username, password, role, is_active ) VALUES($1, $2, $3, $4, $5) RETURNING id"
	INSERT_CUST_USR    = "INSERT INTO ms_customer (id, id_user, full_name, NIK, noPhone, email, address, created_at, created_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);"
	GET_CUST_ID        = "SELECT id,id_user,full_name,NIK,noPhone,email,address,created_at,updated_at,created_by,updated_by FROM ms_customer WHERE id=$1"
	GET_CUST_ID_MEMBER = "SELECT msm.type FROM ms_customer AS msc JOIN ms_member AS msm ON msc.id = msm.id_customer WHERE msc.id = $1 AND msm.expire > CURRENT_DATE"
	GET_CUST_USRID     = "SELECT id,id_user,full_name,NIK,noPhone,email,address,created_at,updated_at,created_by,updated_by FROM ms_customer WHERE id_user = $1"
	GET_ALL_CUSTOMER   = "SELECT id,id_user,full_name,NIK,noPhone,email,address,created_at,updated_at,created_by,updated_by FROM ms_customer"
	GET_CUST_NAME      = "SELECT id,id_user,full_name,NIK,noPhone,email,address,created_at,updated_at,created_by,updated_by FROM ms_customer WHERE full_name = $1"
	EDIT_CUST_ID       = "UPDATE ms_customer SET full_name=$1,NIK=$2,noPhone=$3,email=$4,address=$5,updated_at=$6,updated_by=$7 WHERE id = $8"

	//MEMBER
	INSERT_MEMBER  = "insert into ms_member(id, id_customer, type, expire, created_at, created_by) values($1,$2,$3,$4,$5,$6)"
	GET_ALL_MEMBER = "select id, id_customer, type, expire, created_at, created_by, updated_at, updated_by from ms_member"
	GET_MEMBER_ID  = "SELECT id, id_customer, type, expire, created_at, created_by, updated_at, updated_by from ms_member WHERE id =$1"
	EDIT_MEMBER    = "UPDATE ms_member SET updated_at=$1, updated_by=$2, type=$3, expire=$4 WHERE id=$5"
	DELETE_MEMBER  = "DELETE FROM ms_member WHERE id=$1"

	//RENT
	INSERT_RENT_S = "insert into tx_rent(id, id_vehicle, id_customer, price, date_out, status, created_by,duration) values($1,$2,$3,$4,$5,$6,$7,$8)"
	INSERT_RENT_V = "Update ms_vehicle SET is_available=$1 WHERE id =$2"
	GET_ALL_RENT  = "select id, id_vehicle, id_customer, price, date_in, date_out, status,created_by, updated_by, duration from tx_rent"
	GET_RENT_ID   = "select id, id_vehicle, id_customer, price, date_in, date_out, status,created_by, updated_by, duration from tx_rent where id=$1"
	EDIT_RENT_S   = "UPDATE tx_rent SET date_in =$1, status=$2, price=$3 ,updated_by=$4 WHERE id=$5"
	EDIT_RENT_V   = "Update ms_vehicle SET is_available=$1 WHERE id =$2"

	// CREDIT
	INSERT_CREDIT                = "insert into tx_credit(id, id_vehicle, id_customer, price, interest, date_out, created_at, created_by, credit_duration) values($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id"
	INSERT_INSTALLMENT_CREDIT    = "insert into tx_installment_credit(id, id_vehicle, id_credit, price, total_payment_now, date_payment, date_finish, due_date, status, suspend, current_credit) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) "
	UPDATE_VEHICLE_CREDIT        = "Update ms_vehicle SET is_available=$1 WHERE id =$2"
	GET_ALL_CREDIT               = "select c.id, c.id_vehicle, c.id_customer, c.price, c.interest, c.date_in , c.date_out, c.created_at, c.created_by, c.updated_at, c.credit_duration, c.updated_by, dc.id, dc.id_vehicle, dc.id_credit, dc.price, dc.total_payment_now, dc.date_payment, dc.date_finish, dc.due_date, dc.status, dc.suspend, dc.current_credit from tx_credit c JOIN tx_installment_credit dc ON c.id=dc.id_credit"
	GET_ALL_STRIKE_CREDIT        = "SELECT tc.id, tc.id_customer, tc.id_vehicle, tc.price, tc.interest,tc.date_in, tc.date_out, tc.credit_duration, tic.price, tic.total_payment_now, tic.date_payment, tic.due_date, tic.current_credit, tic.suspend FROM tx_credit tc JOIN tx_installment_credit tic ON tc.id = tic.id_credit WHERE tic.suspend = true  AND tic.status = false  AND tc.date_in = '0001-01-01 00:00:00.000'::timestamp;"
	GET_CREDIT_BY_ID             = "select c.id, c.id_vehicle, c.id_customer, c.price, c.interest, c.date_in , c.date_out, c.created_at, c.created_by, c.updated_at, c.updated_by, c.credit_duration, dc.id, dc.id_vehicle, dc.id_credit, dc.price, dc.total_payment_now, dc.date_payment, dc.date_finish, dc.due_date, dc.status, dc.suspend, dc.current_credit from tx_credit c JOIN tx_installment_credit dc ON c.id=dc.id_credit where c.id=$1"
	GET_INSTALLMENT_CREDIT_BY_ID = "select c.credit_duration, c.Interest, dc.price, dc.total_payment_now, dc.current_credit from tx_credit c JOIN tx_installment_credit dc ON c.id=dc.id_credit where c.id=$1"
	EDIT_CREDIT                  = "UPDATE tx_credit SET date_in=$1, updated_at=$2, updated_by=$3 WHERE id=$4"
	EDIT_INSTALLMENT_CREDIT      = "UPDATE tx_installment_credit SET total_payment_now=$1, date_payment=$2, date_finish=$3, status=$4, suspend=$5, current_credit=$6 WHERE id_credit=$7"
	EDIT_VEHICLE_CREDIT          = "Update ms_vehicle SET is_available=$1 WHERE id =$2"
	GET_SUSPEND_VEHICLE          = "select  tc.id_vehicle from tx_credit tc join tx_installment_credit tic on tc.id = tic.id_credit where tic.suspend = true and tc.id = $1"
	UPDATE_VEHICLE_STATUS        = "update ms_vehicle set is_available=true where id = $1"
	UPDATE_DATEIN_CREDIT         = "update tx_credit set date_in=$1,updated_by=$2,updated_at=$3 where id=$4"

	// Scheduler
	GET_ALL_STRIKE        = "SELECT tic.id, tic.id_vehicle, tic.id_credit, tc.interest FROM tx_installment_credit as tic join tx_credit tc on tic.id_credit = tc.id where status = true and  date_payment <= (CURRENT_DATE - INTERVAL '3 months');"
	UPDATE_STATUS_SUSPEND = "update tx_installment_credit set status=$1, suspend=$2 where id = $3"
	UPDATE_INTEREST       = "update tx_credit set interest=$1 where id=$2"
)
