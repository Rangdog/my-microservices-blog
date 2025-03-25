package discovery

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/consul/api"
)

type ConsulClient struct {
	client    *api.Client
	serviceID string
}

func NewConsulClient(consulAddr, serviceName, host string, port int) (*ConsulClient, error) {
	config := api.DefaultConfig()
	config.Address = consulAddr
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	serviceID := fmt.Sprintf("%s-%s-%d", serviceName, host, port)
	reg := &api.AgentServiceRegistration{
		ID: serviceID,
		Name: serviceName,
		Address: host,
		Port: port,
		Check: &api.AgentServiceCheck{
			HTTP: fmt.Sprintf("http://%s:%d/health", host, port),
			Interval: "10s",
			Timeout: "1s",
		},
	}

	if err := client.Agent().ServiceRegister(reg); err != nil{
		return nil, err
	}
	return &ConsulClient{client: client, serviceID: serviceID}, nil
}

func (c *ConsulClient) StartHeartbeat(ctx context.Context){
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	
	for{
		select{
		case <- ctx.Done():
			c.Deregister()
			return
		case <- ticker.C:
			if err := c.client.Agent().PassTTL("service:" + c.serviceID, ""); err != nil{
				log.Printf("Failed to send heartbeat: %v", err)
			}
		}
	}
}

func (c *ConsulClient) Deregister(){
	if err := c.client.Agent().ServiceDeregister(c.serviceID); err != nil{
		log.Printf("Failed to deregister service: %v", err)
	}
}

func (c *ConsulClient) DiscoverService(serviceName string) ([]*api.ServiceEntry, error){
	entries, _, err := c.client.Health().Service(serviceName, "", true, nil)
	if err != nil{
		return nil, err
	}
	return entries, err
}