package parser

import (
	"encoding/json"
)

type Services struct {
	Redis struct {
		Port     string
		Hostname string
		Password string
	}
}

func (s *Services) Parse(service string) {
	var serv interface{}

	err := json.Unmarshal([]byte(service), &serv)
	if err != nil {
		// This should return err if there are no services attatched and we should relay this message back to the ui and exit
		panic(err)
	}

	map_serv := serv.(map[string]interface{})

	if map_serv["user-provided"] != nil {
		s.userDefinedServices(map_serv["user-provided"])
	}

	//TODO: bosh provided services
}

func (s *Services) userDefinedServices(m interface{}) {
	//TODO parse a user defined service for redis
	m_services := m.([]interface{})
	for _, value := range m_services {
		inner_value := value.(map[string]interface{})

		if inner_value["name"] == "performance-test-redis" {
			credentials := inner_value["credentials"]
			inner_credentials := credentials.(map[string]interface{})

			if inner_credentials["port"] != nil && inner_credentials["hostname"] != nil {
				s.Redis.Port = inner_credentials["port"].(string)
				s.Redis.Hostname = inner_credentials["hostname"].(string)
			} else {
				//TODO: send an error message to the UI if the port and hostname are not properly configured
			}
			if inner_credentials["password"] != nil {
				s.Redis.Password = inner_credentials["password"].(string)
			} else {
				s.Redis.Password = ""
			}
		}
	}
}
