package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Customer struct {
	id        int32
	name      string
	status    bool
	createdAt time.Time
}
type Driver struct {
	id        int32
	name      string
	status    bool
	createdAt time.Time
}
type Order struct {
	id         int32
	driverId   int32
	customerId int32
	daerah     string
	createdAt  time.Time
}

func main() {
	connStr := "user=postgres dbname=Ojek sslmode=disable password=superUser host=localhost"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Mengatur parameter pool koneksi
	db.SetMaxOpenConns(25)                  // Jumlah maksimal koneksi terbuka
	db.SetMaxIdleConns(25)                  // Jumlah maksimal koneksi idle
	db.SetConnMaxLifetime(30 * time.Minute) // Durasi maksimal koneksi (misalnya, 30 menit)
	db.SetConnMaxIdleTime(5 * time.Minute)  // Durasi maksimal koneksi idle (misalnya, 5 menit)

	query := `select count("createdAt") as "total order tiap bulan" from "Order" o  where date_part ('month', o."createdAt") = 10;`
	row := db.QueryRow(query)
	output := struct {
		totalOrderPerMonth int
	}{}
	err = row.Scan(&output.totalOrderPerMonth)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No order found with the provided credentials.")
		} else {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("order found: %+v\n", output)
	}

	query2 :=
		`select c."name", t."total_orders"
	from "Customer" c
	join (
	SELECT o."customer_id",
    	count(*) AS total_orders
	FROM "Order" o
	WHERE 
    	date_part ('month', o."createdAt") = 10
    	and date_part ('year', o."createdAt") = 2024
	GROUP BY 
    	o."customer_id" 
	ORDER BY 
    	total_orders desc) t on c."id" = t."customer_id";`
	row2 := db.QueryRow(query2)
	output2 := struct {
		name       string
		totalOrder int
	}{}
	err = row2.Scan(&output2.name, &output2.totalOrder)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No order found with the provided credentials.")
		} else {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("customer favorite: %+v\n", output2)
	}

	query3 :=
		`select  o."daerah",
	count(*) AS total_orders
	from "Order" o
	where  
	date_part ('month', o."createdAt") = 10
	and date_part ('year', o."createdAt") = 2024
	group  by  
		o."daerah" 
	order  by 
		total_orders desc`
	row3 := db.QueryRow(query3)
	output3 := struct {
		daerah     string
		totalOrder int
	}{}
	err = row3.Scan(&output3.daerah, &output3.totalOrder)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No order found with the provided credentials.")
		} else {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("daerah favorite: %+v\n", output3)
	}

	query4 :=
		`select count(c."status") as "customer online",
	(select count(c."status") as "customer offline"
	from "Customer" c
	where c."status" = false)
	from "Customer" c
	where c."status" = true;`
	row4 := db.QueryRow(query4)
	output4 := struct {
		CustomerOnline  int
		CustomerOffline int
	}{}
	err = row4.Scan(&output4.CustomerOnline, &output4.CustomerOffline)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No order found with the provided credentials.")
		} else {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("Customer Status: %+v\n", output4)
	}

	query5 :=
		`select d."name", t."total_orders"
	from "Driver" d
	join (
	SELECT o."driver_id",
    	COUNT(*) AS total_orders
	FROM "Order" o
	WHERE 
    	date_part ('month', o."createdAt") = 10
    	and date_part ('year', o."createdAt") = 2024
	GROUP BY 
    	o."driver_id" 
	ORDER BY 
    	total_orders desc) t on d."id" = t."driver_id";`
	row5 := db.QueryRow(query5)
	output5 := struct {
		name            string
		totalOrderTaken int
	}{}
	err = row5.Scan(&output5.name, &output5.totalOrderTaken)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No order found with the provided credentials.")
		} else {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("driver favorite: %+v\n", output5)
	}
	query6 :=
		`select date_part ('hour', o."createdAt") as "jam order",
	count(*) AS "total orders"
	from "Order" o
	where  
	date_part ('month', o."createdAt") = 10
	and date_part ('year', o."createdAt") = 2024
	group  by 
		date_part ('hour', o."createdAt")
	order  by 
		"total orders" desc;`
	row6 := db.QueryRow(query6)
	output6 := struct {
		jam        int
		totalOrder int
	}{}
	err = row6.Scan(&output6.jam, &output6.totalOrder)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No order found with the provided credentials.")
		} else {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("prime time order: %+v\n", output6)
	}
}
