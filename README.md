# testgolang
test golang and graphql

test.postman_collection.json

{
	"info": {
		"_postman_id": "ce21e0b1-93b3-4a21-a2ae-0cec077ae7d6",
		"name": "test",
		"schema": "https://schema.getpostman.com/json/collection/v2.0.0/collection.json"
	},
	"item": [
		{
			"name": "localhost:5000/admin/login",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\"username\":\"root@root.com\",\"password\":\"123456\"}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": "localhost:5000/admin/login"
			},
			"response": []
		},
		{
			"name": "localhost:5000/graphl_public",
			"request": {
				"auth": {
					"type": "bearer",
					"bearer": {
						"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0eXBlcyI6ImFkbWluIiwidXNlcm5hbWUiOiJyb290QHJvb3QuY29tIn0.Pnp7ENCvPSiIO5cSAW9rd7qk0HFWZdgz5XFQr1skLp8"
					}
				},
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:5000/graphql_public?query=mutation+_{createclient(email:\"cliente4@gmail.com\",password:\"800800\"){id,email,password}}",
					"host": [
						"localhost"
					],
					"port": "5000",
					"path": [
						"graphql_public"
					],
					"query": [
						{
							"key": "query",
							"value": "mutation+_{createclient(email:\"cliente4@gmail.com\",password:\"800800\"){id,email,password}}"
						}
					]
				}
			},
			"response": []
		},
		{
			"name": "localhost:5000/users/login",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\"username\":\"cliente4@gmail.com\",\"password\":\"800800\"}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": "localhost:5000/users/login"
			},
			"response": []
		},
		{
			"name": "localhost:5000/graphql?query=mutation+_{createtehnical(email:\"tecnico8@gmail.com\",password:\"800800\"){id,email,password}}",
			"request": {
				"auth": {
					"type": "bearer",
					"bearer": {
						"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0eXBlcyI6ImFkbWluIiwidXNlcm5hbWUiOiJyb290QHJvb3QuY29tIn0.Pnp7ENCvPSiIO5cSAW9rd7qk0HFWZdgz5XFQr1skLp8"
					}
				},
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:5000/graphql?query=mutation+_{createtehnical(email:\"tecnico8@gmail.com\",password:\"800800\"){id,email,password}}",
					"host": [
						"localhost"
					],
					"port": "5000",
					"path": [
						"graphql"
					],
					"query": [
						{
							"key": "query",
							"value": "mutation+_{createtehnical(email:\"tecnico8@gmail.com\",password:\"800800\"){id,email,password}}"
						}
					]
				}
			},
			"response": []
		},
		{
			"name": "localhost:5000/graphql?query=mutation+_{updatetehnical(email:\"tecnico8@gmail.com\",password:\"2991212\"){email,password}}",
			"request": {
				"auth": {
					"type": "bearer",
					"bearer": {
						"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0eXBlcyI6ImFkbWluIiwidXNlcm5hbWUiOiJyb290QHJvb3QuY29tIn0.Pnp7ENCvPSiIO5cSAW9rd7qk0HFWZdgz5XFQr1skLp8"
					}
				},
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:5000/graphql?query=mutation+_{updatetehnical(email:\"tecnico8@gmail.com\",password:\"2991212\"){email,password}}",
					"host": [
						"localhost"
					],
					"port": "5000",
					"path": [
						"graphql"
					],
					"query": [
						{
							"key": "query",
							"value": "mutation+_{updatetehnical(email:\"tecnico8@gmail.com\",password:\"2991212\"){email,password}}"
						}
					]
				}
			},
			"response": []
		},
		{
			"name": "localhost:5000/graphql?query=mutation+_{deletetehnical(email:\"tecnico8@gmail.com\"){id,email}}",
			"request": {
				"auth": {
					"type": "bearer",
					"bearer": {
						"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0eXBlcyI6ImFkbWluIiwidXNlcm5hbWUiOiJyb290QHJvb3QuY29tIn0.Pnp7ENCvPSiIO5cSAW9rd7qk0HFWZdgz5XFQr1skLp8"
					}
				},
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:5000/graphql?query=mutation+_{deletetehnical(email:\"tecnico8@gmail.com\"){id,email}}",
					"host": [
						"localhost"
					],
					"port": "5000",
					"path": [
						"graphql"
					],
					"query": [
						{
							"key": "query",
							"value": "mutation+_{deletetehnical(email:\"tecnico8@gmail.com\"){id,email}}"
						}
					]
				}
			},
			"response": []
		},
		{
			"name": "localhost:5000/graphql?query=mutation+_{updateclient(email:\"cliente3@gmail.com\",password:\"20202\"){email,password}}",
			"request": {
				"auth": {
					"type": "bearer",
					"bearer": {
						"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0eXBlcyI6ImFkbWluIiwidXNlcm5hbWUiOiJyb290QHJvb3QuY29tIn0.Pnp7ENCvPSiIO5cSAW9rd7qk0HFWZdgz5XFQr1skLp8"
					}
				},
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:5000/graphql?query=mutation+_{updateclient(email:\"cliente3@gmail.com\",password:\"20202\"){email,password}}",
					"host": [
						"localhost"
					],
					"port": "5000",
					"path": [
						"graphql"
					],
					"query": [
						{
							"key": "query",
							"value": "mutation+_{updateclient(email:\"cliente3@gmail.com\",password:\"20202\"){email,password}}"
						}
					]
				}
			},
			"response": []
		},
		{
			"name": "localhost:5000/graphql?query=mutation+_{deleteclient(email:\"cliente3@gmail.com\"){id,email,password}}",
			"request": {
				"auth": {
					"type": "bearer",
					"bearer": {
						"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0eXBlcyI6ImFkbWluIiwidXNlcm5hbWUiOiJyb290QHJvb3QuY29tIn0.Pnp7ENCvPSiIO5cSAW9rd7qk0HFWZdgz5XFQr1skLp8"
					}
				},
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:5000/graphql?query=mutation+_{deleteclient(email:\"cliente3@gmail.com\"){id,email,password}}",
					"host": [
						"localhost"
					],
					"port": "5000",
					"path": [
						"graphql"
					],
					"query": [
						{
							"key": "query",
							"value": "mutation+_{deleteclient(email:\"cliente3@gmail.com\"){id,email,password}}"
						}
					]
				}
			},
			"response": []
		},
		{
			"name": "localhost:5000/graphql?query=mutation+_{createticket(typeticket:\"Install\",dayinsert:\"18-09-2020\",dayasign:\"19-09-2020\",idclient:53055,status:\"init\",address:\"Barranquilla cra 7d calle 35a\"){id,typeticket,address}}",
			"request": {
				"auth": {
					"type": "bearer",
					"bearer": {
						"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0eXBlcyI6ImFkbWluIiwidXNlcm5hbWUiOiJyb290QHJvb3QuY29tIn0.Pnp7ENCvPSiIO5cSAW9rd7qk0HFWZdgz5XFQr1skLp8"
					}
				},
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:5000/graphql?query=mutation+_{createticket(typeticket:\"Install\",dayinsert:\"18-09-2020\",dayasign:\"19-09-2020\",idclient:53055,status:\"init\",address:\"Barranquilla cra 7d calle 35a\"){id,typeticket,address}}",
					"host": [
						"localhost"
					],
					"port": "5000",
					"path": [
						"graphql"
					],
					"query": [
						{
							"key": "query",
							"value": "mutation+_{createticket(typeticket:\"Install\",dayinsert:\"18-09-2020\",dayasign:\"19-09-2020\",idclient:53055,status:\"init\",address:\"Barranquilla cra 7d calle 35a\"){id,typeticket,address}}"
						}
					]
				}
			},
			"response": []
		},
		{
			"name": "localhost:5000/graphql?query=mutation+_{updateticket(id:64733,status:\"ok\"){id,status}}",
			"request": {
				"auth": {
					"type": "bearer",
					"bearer": {
						"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0eXBlcyI6ImFkbWluIiwidXNlcm5hbWUiOiJyb290QHJvb3QuY29tIn0.Pnp7ENCvPSiIO5cSAW9rd7qk0HFWZdgz5XFQr1skLp8"
					}
				},
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:5000/graphql?query=mutation+_{updateticket(id:64733,status:\"ok\"){id,status}}",
					"host": [
						"localhost"
					],
					"port": "5000",
					"path": [
						"graphql"
					],
					"query": [
						{
							"key": "query",
							"value": "mutation+_{updateticket(id:64733,status:\"ok\"){id,status}}"
						}
					]
				}
			},
			"response": []
		},
		{
			"name": "localhost:5000/graphql?query=mutation+_{deleteticket(id:4404){id,status}}",
			"request": {
				"auth": {
					"type": "bearer",
					"bearer": {
						"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0eXBlcyI6ImFkbWluIiwidXNlcm5hbWUiOiJyb290QHJvb3QuY29tIn0.Pnp7ENCvPSiIO5cSAW9rd7qk0HFWZdgz5XFQr1skLp8"
					}
				},
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:5000/graphql?query=mutation+_{deleteticket(id:4404){id,status}}",
					"host": [
						"localhost"
					],
					"port": "5000",
					"path": [
						"graphql"
					],
					"query": [
						{
							"key": "query",
							"value": "mutation+_{deleteticket(id:4404){id,status}}"
						}
					]
				}
			},
			"response": []
		},
		{
			"name": "localhost:5000/graphql?query=mutation+_{listday(idtechnical:53954){id,status}}",
			"request": {
				"auth": {
					"type": "bearer",
					"bearer": {
						"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0eXBlcyI6ImFkbWluIiwidXNlcm5hbWUiOiJyb290QHJvb3QuY29tIn0.Pnp7ENCvPSiIO5cSAW9rd7qk0HFWZdgz5XFQr1skLp8"
					}
				},
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:5000/graphql?query=mutation+_{listday(idtechnical:53954){id,status}}",
					"host": [
						"localhost"
					],
					"port": "5000",
					"path": [
						"graphql"
					],
					"query": [
						{
							"key": "query",
							"value": "mutation+_{listday(idtechnical:53954){id,status}}"
						}
					]
				}
			},
			"response": []
		},
		{
			"name": "localhost:5000/graphql?query={technical(email:\"tecnico1@gmail.com\"){email}}",
			"request": {
				"auth": {
					"type": "bearer",
					"bearer": {
						"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0eXBlcyI6ImFkbWluIiwidXNlcm5hbWUiOiJyb290QHJvb3QuY29tIn0.Pnp7ENCvPSiIO5cSAW9rd7qk0HFWZdgz5XFQr1skLp8"
					}
				},
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:5000/graphql?query={technical(email:\"tecnico1@gmail.com\"){email}}",
					"host": [
						"localhost"
					],
					"port": "5000",
					"path": [
						"graphql"
					],
					"query": [
						{
							"key": "query",
							"value": "{technical(email:\"tecnico1@gmail.com\"){email}}"
						}
					]
				}
			},
			"response": []
		},
		{
			"name": "localhost:5000/graphql?query={client(email:\"cliente4@gmail.com\"){email}}",
			"request": {
				"auth": {
					"type": "bearer",
					"bearer": {
						"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0eXBlcyI6ImFkbWluIiwidXNlcm5hbWUiOiJyb290QHJvb3QuY29tIn0.Pnp7ENCvPSiIO5cSAW9rd7qk0HFWZdgz5XFQr1skLp8"
					}
				},
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:5000/graphql?query={client(email:\"cliente4@gmail.com\"){email}}",
					"host": [
						"localhost"
					],
					"port": "5000",
					"path": [
						"graphql"
					],
					"query": [
						{
							"key": "query",
							"value": "{client(email:\"cliente4@gmail.com\"){email}}"
						}
					]
				}
			},
			"response": []
		},
		{
			"name": "localhost:5000/graphql?query={listtehnicals{id,email,password}}",
			"request": {
				"auth": {
					"type": "bearer",
					"bearer": {
						"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0eXBlcyI6ImFkbWluIiwidXNlcm5hbWUiOiJyb290QHJvb3QuY29tIn0.Pnp7ENCvPSiIO5cSAW9rd7qk0HFWZdgz5XFQr1skLp8"
					}
				},
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:5000/graphql?query={listtehnicals{id,email,password}}",
					"host": [
						"localhost"
					],
					"port": "5000",
					"path": [
						"graphql"
					],
					"query": [
						{
							"key": "query",
							"value": "{listtehnicals{id,email,password}}"
						}
					]
				}
			},
			"response": []
		},
		{
			"name": "localhost:5000/graphql?query={listclients{id,email,password}}",
			"request": {
				"auth": {
					"type": "bearer",
					"bearer": {
						"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0eXBlcyI6ImFkbWluIiwidXNlcm5hbWUiOiJyb290QHJvb3QuY29tIn0.Pnp7ENCvPSiIO5cSAW9rd7qk0HFWZdgz5XFQr1skLp8"
					}
				},
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:5000/graphql?query={listclients{id,email,password}}",
					"host": [
						"localhost"
					],
					"port": "5000",
					"path": [
						"graphql"
					],
					"query": [
						{
							"key": "query",
							"value": "{listclients{id,email,password}}"
						}
					]
				}
			},
			"response": []
		},
		{
			"name": "localhost:5000/graphql?query={listtickets{id,status,dayasign}}",
			"request": {
				"auth": {
					"type": "bearer",
					"bearer": {
						"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0eXBlcyI6ImFkbWluIiwidXNlcm5hbWUiOiJyb290QHJvb3QuY29tIn0.Pnp7ENCvPSiIO5cSAW9rd7qk0HFWZdgz5XFQr1skLp8"
					}
				},
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:5000/graphql?query={listtickets{id,status,dayasign}}",
					"host": [
						"localhost"
					],
					"port": "5000",
					"path": [
						"graphql"
					],
					"query": [
						{
							"key": "query",
							"value": "{listtickets{id,status,dayasign}}"
						}
					]
				}
			},
			"response": []
		}
	],
	"protocolProfileBehavior": {}
}
