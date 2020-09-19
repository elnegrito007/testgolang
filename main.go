package main

import (
	"encoding/json"
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
	"github.com/go-redis/redis"
	"github.com/graphql-go/graphql"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	master "test/env"
	shorts "test/functions"
	admin "test/shemas"
	"time"
)

var tokenAuth *jwtauth.JWTAuth

var client = redis.NewClient(&redis.Options{
	Addr:     master.Host(),
	Password: master.Password(),
	DB:       0,
})

func init() {
	tokenAuth = jwtauth.New("HS256", []byte(master.KeySha()), nil)
	_ = client.Set("admin_root@root.com", `{"id":1000,"email":"root@root.com","password":"123456"}`, 0).Err()
}

func main() {
	fmt.Println("Server on port "+master.Port(), time.Now().UTC())
	_ = http.ListenAndServe(master.Port(), router())
}

func AdminLogin(w http.ResponseWriter, r *http.Request) {
	var user admin.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	regCorreo, regPass := regexp.MustCompile(master.EmailReg()), regexp.MustCompile(master.PassReg())
	if user.Username == "" || user.Password == "" || regCorreo.MatchString(user.Username) == false || regPass.MatchString(user.Password) == false {
		_, _ = fmt.Fprintf(w, "{\"e\":1,\"d\":[1,2]}")
	} else {
		val, err := client.Get("admin_" + user.Username).Result()
		if err == nil {
			data := []byte(val)
			var clave, _, _, _ = jsonparser.Get(data, "password")
			if string(clave) == user.Password {
				_, tokenString, _ := tokenAuth.Encode(jwt.MapClaims{"username": user.Username,"types":"admin"})
				w.Header().Set("content-type", "application/json")
				_, _ = w.Write([]byte(`{ "token": "` + tokenString + `" }`))
			}else{
				_, _ = fmt.Fprintf(w, "{\"e\":1,\"d\":[1,3]}")
			}
		} else {
			_, _ = fmt.Fprintf(w, "{\"e\":1,\"d\":[1,1]}")
		}
	}
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	var user admin.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	regCorreo, regPass := regexp.MustCompile(master.EmailReg()), regexp.MustCompile(master.PassReg())
	if user.Username == "" || user.Password == "" || regCorreo.MatchString(user.Username) == false || regPass.MatchString(user.Password) == false {
		_, _ = fmt.Fprintf(w, "{\"e\":1,\"d\":[1,2]}")
	} else {
		val, err := client.Get("users_" + user.Username).Result()
		if err == nil {
			data := []byte(val)
			var clave, _, _, _ = jsonparser.Get(data, "password")
			if string(clave) == user.Password {
				_, tokenString, _ := tokenAuth.Encode(jwt.MapClaims{"username": user.Username,"types":"users"})
				w.Header().Set("content-type", "application/json")
				_, _ = w.Write([]byte(`{ "token": "` + tokenString + `" }`))
			}else{
				_, _ = fmt.Fprintf(w, "{\"e\":1,\"d\":[1,3]}")
			}
		} else {
			_, _ = fmt.Fprintf(w, "{\"e\":1,\"d\":[1,1]}")
		}
	}
}

func router() http.Handler {

	r := chi.NewRouter()

	corsVar := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	r.Use(corsVar.Handler)

	type Technical struct {
		ID    int   `json:"id"`
		Email  string  `json:"email"`
		Password  string  `json:"password"`
	}

	type Ticket struct {
		ID    int   `json:"id"`
		Typeticket  string  `json:"typeticket"`
		DayInsert  string  `json:"dayinsert"`
		DayAsign  string  `json:"dayasign"`
		IdClient  int  `json:"idclient"`
		IdTechnical  int  `json:"idtechnical"`
		Status  string  `json:"status"`
		Address  string  `json:"address"`
	}

	type Client struct {
		ID    int   `json:"id"`
		Email  string  `json:"email"`
		Password  string  `json:"password"`
	}

	var technicals []Technical
	var keys []string
	var cursor uint64
	keys, _, _ = client.Scan(cursor, "admin_*", 100).Result()
	sum := 0
	for sum < len(keys) {
		val, _ := client.Get(keys[sum]).Result()
		sum++
		var app Technical
		_= json.Unmarshal([]byte(val), &app)
		technicals = append(technicals,app)
	}
	fmt.Println(technicals)

	var clients []Client
	var keys2 []string
	var cursor2 uint64
	keys2, _, _ = client.Scan(cursor2, "client_*", 100).Result()
	sum2 := 0
	for sum2 < len(keys2) {
		val, _ := client.Get(keys2[sum2]).Result()
		sum2++
		var app Client
		_= json.Unmarshal([]byte(val), &app)
		clients = append(clients,app)
	}
	fmt.Println(clients)

	var Tickets []Ticket
	var keys3 []string
	var cursor3 uint64
	keys3, _, _ = client.Scan(cursor3, "ticket_*", 100).Result()
	sum3 := 0
	for sum3 < len(keys3) {
		val, _ := client.Get(keys3[sum3]).Result()
		sum3++
		var app Ticket
		_= json.Unmarshal([]byte(val), &app)
		Tickets = append(Tickets,app)
	}
	fmt.Println(Tickets)

	var technicalType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "technical",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"email": &graphql.Field{
					Type: graphql.String,
				},
				"password": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)

	var ticketType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "ticket",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"typeticket": &graphql.Field{
					Type: graphql.String,
				},
				"dayinsert": &graphql.Field{
					Type: graphql.String,
				},
				"dayasign": &graphql.Field{
					Type: graphql.String,
				},
				"idclient": &graphql.Field{
					Type: graphql.Int,
				},
				"idtechnical": &graphql.Field{
					Type: graphql.Int,
				},
				"status": &graphql.Field{
					Type: graphql.String,
				},
				"address": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)

	var clientType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "client",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"email": &graphql.Field{
					Type: graphql.String,
				},
				"password": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)

	var mutationTypePublic = graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{

			/* Create new client item
			localhost:5000/graphql_public?query=mutation+_{createclient(email:"prueba",password:"800800"){id,email,password}}
			*/
			"createclient": &graphql.Field{
				Type:        clientType,
				Description: "Create new client",
				Args: graphql.FieldConfigArgument{
					"email": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					rand.Seed(time.Now().UnixNano())
					clients := Client{
						ID:       rand.Intn(100000),
						Email:    params.Args["email"].(string),
						Password: params.Args["password"].(string),
					}
					_ = client.Set("client_"+params.Args["email"].(string), `{"email":"`+params.Args["email"].(string)+`","password":"`+params.Args["password"].(string)+`","id":`+strconv.Itoa(clients.ID)+`}`, 0).Err()
					return clients, nil
				},
			},

		},
	})

	var mutationType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{

			/* Create new tehnical item
			localhost:5000/graphql?query=mutation+_{createtehnical(email:"tecnico1@gmail.com",password:"800800"){id,email,password}}
			*/
			"createtehnical": &graphql.Field{
				Type:        technicalType,
				Description: "Create new technical",
				Args: graphql.FieldConfigArgument{
					"email": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					rand.Seed(time.Now().UnixNano())
					technical := Technical{
						ID:    rand.Intn(100000),
						Email:  params.Args["email"].(string),
						Password:  params.Args["password"].(string),
					}
					technicals =  append(technicals,technical)
					valx, _ := client.Get("admin_"+params.Args["email"].(string)).Result()
					if valx == "" {
						_ = client.Set("admin_"+params.Args["email"].(string), `{"email":"`+params.Args["email"].(string)+`","password":"`+params.Args["password"].(string)+`","id":`+strconv.Itoa(technical.ID)+`}`, 0).Err()
						return technical, nil
					}else{
						return nil, nil
					}
				},
			},

			/* Update tehnical by id
			   localhost:5000/graphql?query=mutation+_{updatetehnical(email:"tecnico1@gmail.com",password:"2991212"){email,password}}
			*/
			"updatetehnical": &graphql.Field{
				Type:        technicalType,
				Description: "Update technical by id",
				Args: graphql.FieldConfigArgument{
					"email": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					email, emailOk  := params.Args["email"].(string)
					password, passwordOk := params.Args["password"].(string)
					technical := Technical{
						ID:    0,
						Email:  params.Args["email"].(string),
						Password:  params.Args["password"].(string),
					}
					if passwordOk && emailOk {
						val, err := client.Get("admin_"+email).Result()
						if err == nil {
							data := []byte(val)
							var id, _, _, _= jsonparser.Get(data, "id")
							_ = client.Set("admin_"+email, `{"email":"`+email+`","password":"`+password+`","id":`+string(id)+`}`, 0).Err()
						}
					}
					return technical, nil
				},
			},

			/* Delete tehnical by id
			   localhost:5000/graphql?query=mutation+_{deletetehnical(email:"tecnico1@gmail.com"){id,email}}
			*/
			"deletetehnical": &graphql.Field{
				Type:        technicalType,
				Description: "Delete technical by email",
				Args: graphql.FieldConfigArgument{
					"email": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					email, _ := params.Args["email"].(string)
					technical := Technical{}
					for i, p := range technicals {
						if email == p.Email {
							technical = technicals[i]
							technicals = append(technicals[:i], technicals[i+1:]...)
						}
					}
					_ = client.Del("admin_"+email).Err()
					return technical, nil
				},
			},

			/* Update client by id
			   localhost:5000/graphql?query=mutation+_{updateclient(email:"cliente4@gmail.com",password:"20202"){email,password}}
			*/
			"updateclient": &graphql.Field{
				Type:        clientType,
				Description: "Update client by id",
				Args: graphql.FieldConfigArgument{
					"email": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					email, emailOk  := params.Args["email"].(string)
					password, passwordOk := params.Args["password"].(string)
					clientAux := Client{
						ID:    0,
						Email:  params.Args["email"].(string),
						Password:  params.Args["password"].(string),
					}
					if passwordOk && emailOk {
						val, err := client.Get("client_"+email).Result()
						if err == nil {
							data := []byte(val)

							for _, p := range clients {
								if email == p.Email {
									p.Password = password
								}
							}

							var id, _, _, _= jsonparser.Get(data, "id")
							_ = client.Set("client_"+email, `{"email":"`+email+`","password":"`+password+`","id":`+string(id)+`}`, 0).Err()
						}
					}
					return clientAux, nil
				},
			},

			/* Delete client by id
			   localhost:5000/graphql?query=mutation+_{deleteclient(id:27872){id,email,password}}
			*/
			"deleteclient": &graphql.Field{
				Type:        clientType,
				Description: "Delete client by email",
				Args: graphql.FieldConfigArgument{
					"email": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					email, _ := params.Args["email"].(string)
					clientsAux := Client{}
					for i, p := range clients {
						if email == p.Email {
							clientsAux = clients[i]
							clients = append(clients[:i], clients[i+1:]...)
						}
					}
					_ = client.Del("client_"+email).Err()
					return clientsAux, nil
				},
			},

			/* Create new ticket item
			localhost:5000/graphql?query=mutation+_{createticket(typeticket:"Install",dayinsert:"18-09-2020",dayasign:"19-09-2020",idclient:61079,status:"init",address:"Barranquilla cra 7d calle 35a"){id,typeticket,address}}
			*/
			"createticket": &graphql.Field{
				Type:        ticketType,
				Description: "Create new ticket",
				Args: graphql.FieldConfigArgument{
					"typeticket": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"dayinsert": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"dayasign": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"idclient": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"status": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"address": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					rand.Seed(time.Now().UnixNano())
					var random = rand.Intn(len(technicals) - 0) + 0
					register := false
					for _, p := range clients {
						if p.ID == params.Args["idclient"].(int) {
							register = true
						}
					}
					if register {
						ticket := Ticket{
							ID:          rand.Intn(100000),
							Typeticket:  params.Args["typeticket"].(string),
							DayInsert:   params.Args["dayinsert"].(string),
							DayAsign:    params.Args["dayasign"].(string),
							IdClient:    params.Args["idclient"].(int),
							IdTechnical: technicals[random].ID,
							Status:      params.Args["status"].(string),
							Address:     params.Args["address"].(string),
						}
						Tickets =  append(Tickets,ticket)
						_ = client.Set("ticket_"+strconv.Itoa(ticket.ID)+"_"+strconv.Itoa(technicals[random].ID)+"_"+params.Args["dayasign"].(string)+"_"+params.Args["status"].(string), `{"typeticket":"`+params.Args["typeticket"].(string)+`","id":`+strconv.Itoa(ticket.ID)+`,"status":"`+params.Args["status"].(string)+`","idtechnical":`+strconv.Itoa(technicals[random].ID)+`,"dayinsert":"`+params.Args["dayinsert"].(string)+`","address":"`+params.Args["address"].(string)+`","idclient":`+strconv.Itoa(params.Args["idclient"].(int))+`,"dayasign":"`+params.Args["dayasign"].(string)+`"}`, 0).Err()
						return ticket, nil
					}else{
						return nil, nil
					}

				},
			},

			/* Update ticket by id
			   localhost:5000/graphql?query=mutation+_{updateticket(id:26865,status:"ok"){id,status}}
			*/
			"updateticket": &graphql.Field{
				Type:        ticketType,
				Description: "Update ticket by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"status": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, idOk  := params.Args["id"].(int)
					status, statusOk := params.Args["status"].(string)

					if idOk && statusOk{

						update := false
						for _, p := range Tickets {
							if p.ID == id {
								update = true
								p.Status = status
							}
						}

						if update {
							valor, _ := client.Keys("ticket_"+strconv.Itoa(id)+"_*").Result()
							val, err := client.Get(valor[0]).Result()
							_ = client.Del(valor[0]).Err()
							if err == nil {
								data := []byte(val)
								var idTicket, _, _, _= jsonparser.Get(data, "id")
								var dayasign, _, _, _= jsonparser.Get(data, "dayasign")
								var typeticket, _, _, _= jsonparser.Get(data, "typeticket")
								var idtechnical, _, _, _= jsonparser.Get(data, "idtechnical")
								var dayinsert, _, _, _= jsonparser.Get(data, "dayinsert")
								var address, _, _, _= jsonparser.Get(data, "address")
								var idclient, _, _, _= jsonparser.Get(data, "idclient")
								_ = client.Set("ticket_"+string(idTicket)+"_"+string(idtechnical)+"_"+string(dayasign)+"_"+status, `{"typeticket":"`+string(typeticket)+`","id":`+string(idTicket)+`,"status":"`+status+`","idtechnical":`+string(idtechnical)+`,"dayinsert":"`+string(dayinsert)+`","address":"`+string(address)+`","idclient":`+string(idclient)+`,"dayasign":"`+string(dayasign)+`"}`, 0).Err()
							}
							technical := Ticket{
								ID:          id,
								Status:      params.Args["status"].(string),
							}
							return technical, nil
						}else{
							return nil, nil
						}
					}else{
						return nil, nil
					}
				},
			},

			/* Delete ticket by id
			   localhost:5000/graphql?query=mutation+_{deleteticket(id:68833){id,status}}
			*/
			"deleteticket": &graphql.Field{
				Type:        ticketType,
				Description: "Delete ticket by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(int)
					valor, _ := client.Keys("ticket_"+strconv.Itoa(id)+"_*").Result()
					_ = client.Del(valor[0]).Err()
					ticket := Ticket{}
					for i, p := range Tickets {
						if id == p.ID {
							ticket = Tickets[i]
							Tickets = append(Tickets[:i], Tickets[i+1:]...)
						}
					}
					return ticket, nil
				},
			},

			/* Update ticket by id in day
			   localhost:5000/graphql?query=mutation+_{listday(idtechnical:91987){id,status}}
			*/
			"listday": &graphql.Field{
				Type:        ticketType,
				Description: "Update list by id tehnical",
				Args: graphql.FieldConfigArgument{
					"idtechnical": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					idtechnical, idOk  := params.Args["idtechnical"].(int)

					if idOk {

						var Tickets2 []Ticket
						ok := false
						for _, p := range Tickets {
							if p.IdTechnical == idtechnical && p.Status=="ok" {
								ok = true
								Tickets2 = append(Tickets2,p)
							}
						}

						if ok {
							return Tickets2[0] ,nil
						}else{
							return 0, nil
						}

					}else{
						return nil, nil
					}
				},
			},

		},
	})

	var queryType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				/* Get (read) single tehnical by id
				localhost:5000/graphql?query={technical(email:"tecnico1@gmail.com"){email}}
				*/
				"technical": &graphql.Field{
					Type: technicalType,
					Description: "Get technical by email",
					Args: graphql.FieldConfigArgument{
						"email": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						email, ok := p.Args["email"].(string)
						if ok {
							for _, technical := range technicals {
								if technical.Email == email {
									return technical, nil
								}
							}
						}
						return nil, nil
					},
				},

				/* Get (read) single client by email
				localhost:5000/graphql?query={client(email:"cliente4@gmail.com"){email}}
				*/
				"client": &graphql.Field{
					Type: clientType,
					Description: "Get client by email",
					Args: graphql.FieldConfigArgument{
						"email": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						email, ok := p.Args["email"].(string)
						if ok {
							// Find technical
							for _, clients := range clients {
								if clients.Email == email {
									return clients, nil
								}
							}
						}
						return nil, nil
					},
				},

				/* Get (read) tehnical list
				   localhost:5000/graphql?query={listtehnicals{id,email,password}}
				*/
				"listtehnicals": &graphql.Field{
					Type:graphql.NewList(technicalType),
					Description: "Get technical list",
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						return technicals, nil
					},
				},

				/* Get (read) client list
				   localhost:5000/graphql?query={listclients{id,email,password}}
				*/
				"listclients": &graphql.Field{
					Type:graphql.NewList(clientType),
					Description: "Get client list",
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						return clients, nil
					},
				},

				/* Get (read) ticket list
				   localhost:5000/graphql?query={listtickets{id,status,dayasign}}
				*/
				"listtickets": &graphql.Field{
					Type:graphql.NewList(ticketType),
					Description: "Get ticket list",
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						return Tickets, nil
					},
				},

			},
		})

	var queryTypePublic = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{},
		})

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
		Mutation: mutationType,
	})

	schema2, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: queryTypePublic,
		Mutation: mutationTypePublic,
	})

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Get("/graphql", func(w http.ResponseWriter, r *http.Request){
			result := graphql.Do(graphql.Params{
				Schema:schema,
				RequestString: r.URL.Query().Get("query"),
			})
			_ = json.NewEncoder(w).Encode(result)
		})
	})

	r.Group(func(r chi.Router) {
		r.Get("/graphql_public", func(w http.ResponseWriter, r *http.Request){
			result := graphql.Do(graphql.Params{
				Schema:schema2,
				RequestString: r.URL.Query().Get("query"),
			})
			_ = json.NewEncoder(w).Encode(result)
		})
		r.Post("/admin/login", AdminLogin)
		r.Post("/users/login", UserLogin)
	})
	r.NotFound(shorts.NotExist)
	return r
}
