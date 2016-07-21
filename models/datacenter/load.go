package datacenter

import (
	"fmt"
	"sync"

	"github.com/centurylinkcloud/clc-go-cli/base"
)

var (
	ListTimeout = 200
)

func All(cn base.Connection, withComputeLimits, withNetworkLimits, withAvailableOvfs, withLoadBalancers bool) ([]GetRes, error) {
	fetcher := dcListFetcher{
		connection:        cn,
		withComputeLimits: withComputeLimits,
		withNetworkLimits: withNetworkLimits,
		withAvailableOvfs: withAvailableOvfs,
		withLoadBalancers: withLoadBalancers,
	}
	return fetcher.Fetch()
}

type dcListFetcher struct {
	connection        base.Connection
	withComputeLimits bool
	withNetworkLimits bool
	withAvailableOvfs bool
	withLoadBalancers bool
}

func (f *dcListFetcher) Fetch() ([]GetRes, error) {
	list, err := f.fetchList()
	if err != nil {
		return list, err
	}

	if f.withComputeLimits {
		err = f.fetchComputeLimits(list)
		if err != nil {
			return list, err
		}
	}

	if f.withNetworkLimits {
		err = f.fetchNetworkLimits(list)
		if err != nil {
			return list, err
		}
	}

	if f.withAvailableOvfs {
		err = f.fetchAvailableOvfs(list)
		if err != nil {
			return list, err
		}
	}

	if f.withLoadBalancers {
		err = f.fetchLoadBalancers(list)
		if err != nil {
			return list, err
		}
	}

	return list, nil
}

func (f *dcListFetcher) fetchList() ([]GetRes, error) {
	var err error

	datacentersUrl := fmt.Sprintf("%s/v2/datacenters/{accountAlias}", base.URL)
	res := []GetRes{}
	err = f.connection.ExecuteRequest("GET", datacentersUrl, nil, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (f *dcListFetcher) fetchComputeLimits(list []GetRes) error {
	var wg sync.WaitGroup

	wg.Add(len(list))

	for i := 0; i < len(list); i++ {
		item := list[i]
		go func(i int) {
			defer wg.Done()
			singleFetcher := dcSingleFetcher{
				connection: f.connection,
			}
			singleFetcher.FetchComputeLimits(&item)
			list[i] = item
		}(i)
	}

	wg.Wait()

	return nil
}

func (f *dcListFetcher) fetchNetworkLimits(list []GetRes) error {
	var wg sync.WaitGroup

	wg.Add(len(list))

	for i := 0; i < len(list); i++ {
		item := list[i]
		go func(i int) {
			defer wg.Done()
			singleFetcher := dcSingleFetcher{
				connection: f.connection,
			}
			singleFetcher.FetchNetworkLimits(&item)
			list[i] = item
		}(i)
	}

	wg.Wait()

	return nil
}

func (f *dcListFetcher) fetchAvailableOvfs(list []GetRes) error {
	var wg sync.WaitGroup

	wg.Add(len(list))

	for i := 0; i < len(list); i++ {
		item := list[i]
		go func(i int) {
			defer wg.Done()
			singleFetcher := dcSingleFetcher{
				connection: f.connection,
			}
			singleFetcher.FetchAvailableOvfs(&item)
			list[i] = item
		}(i)
	}

	wg.Wait()

	return nil
}

func (f *dcListFetcher) fetchLoadBalancers(list []GetRes) error {
	var wg sync.WaitGroup

	wg.Add(len(list))

	for i := 0; i < len(list); i++ {
		item := list[i]
		go func(i int) {
			defer wg.Done()
			singleFetcher := dcSingleFetcher{
				connection: f.connection,
			}
			singleFetcher.FetchLoadBalancers(&item)
			list[i] = item
		}(i)
	}

	wg.Wait()

	return nil
}

func Get(cn base.Connection, dcId string, withComputeLimits, withNetworkLimits, withAvailableOvfs, withLoadBalancers bool) (GetRes, error) {
	fetcher := dcSingleFetcher{
		connection:        cn,
		withComputeLimits: withComputeLimits,
		withNetworkLimits: withNetworkLimits,
		withAvailableOvfs: withAvailableOvfs,
		withLoadBalancers: withLoadBalancers,
	}
	return fetcher.Fetch(dcId)
}

type dcSingleFetcher struct {
	connection        base.Connection
	withComputeLimits bool
	withNetworkLimits bool
	withAvailableOvfs bool
	withLoadBalancers bool
}

func (f *dcSingleFetcher) Fetch(dcId string) (GetRes, error) {
	res, err := f.fetchDc(dcId)

	if f.withComputeLimits {
		f.FetchComputeLimits(&res)
		if err != nil {
			return res, err
		}
	}

	if f.withNetworkLimits {
		err = f.FetchNetworkLimits(&res)
		if err != nil {
			return res, err
		}
	}

	if f.withAvailableOvfs {
		err = f.FetchAvailableOvfs(&res)
		if err != nil {
			return res, err
		}
	}

	if f.withLoadBalancers {
		err = f.FetchLoadBalancers(&res)
		if err != nil {
			return res, err
		}
	}

	return res, err
}

func (f *dcSingleFetcher) fetchDc(dcId string) (GetRes, error) {
	dcUrl := fmt.Sprintf("%s/v2/datacenters/{accountAlias}/%s", base.URL, dcId)
	res := GetRes{}
	err := f.connection.ExecuteRequest("GET", dcUrl, struct{}{}, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (f *dcSingleFetcher) FetchComputeLimits(item *GetRes) error {
	limitsRelativePath, err := item.Links.Get("computeLimits")
	if err != nil {
		return err
	}

	computeLimitsUrl := fmt.Sprintf("%s%s", base.URL, limitsRelativePath)
	res := &DcComputeLimits{}
	f.connection.ExecuteRequest("GET", computeLimitsUrl, struct{}{}, res)
	item.ComputeLimits = res

	return nil
}

func (f *dcSingleFetcher) FetchNetworkLimits(item *GetRes) error {
	limitsRelativePath, err := item.Links.Get("networkLimits")
	if err != nil {
		return err
	}

	networkLimitsUrl := fmt.Sprintf("%s%s", base.URL, limitsRelativePath)
	res := &DcNetworkLimits{}
	f.connection.ExecuteRequest("GET", networkLimitsUrl, struct{}{}, res)
	item.NetworkLimits = res

	return nil
}

func (f *dcSingleFetcher) FetchAvailableOvfs(item *GetRes) error {
	availableOvfsRelativePath, err := item.Links.Get("availableOvfs")
	if err != nil {
		return err
	}

	availableOvfsUrl := fmt.Sprintf("%s%s", base.URL, availableOvfsRelativePath)
	res := []DcAvailableOVF{}
	f.connection.ExecuteRequest("GET", availableOvfsUrl, struct{}{}, &res)
	item.AvailableOVFs = &res

	return nil
}

func (f *dcSingleFetcher) FetchLoadBalancers(item *GetRes) error {
	loadBalancersRelativePath, err := item.Links.Get("loadBalancers")
	if err != nil {
		return err
	}

	loadBalancersUrl := fmt.Sprintf("%s%s", base.URL, loadBalancersRelativePath)
	res := []DcLoadBalancer{}
	f.connection.ExecuteRequest("GET", loadBalancersUrl, struct{}{}, &res)
	item.LoadBalancers = &res

	return nil
}
