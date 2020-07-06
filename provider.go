// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style license that can be found
// in the LICENSE file.

package dnspodcn

import (
	"context"
	"fmt"
	"strconv"

	"github.com/libdns/libdns"
)

var (
	_ libdns.RecordAppender = (*Provider)(nil)
	_ libdns.RecordDeleter  = (*Provider)(nil)
	_ libdns.RecordGetter   = (*Provider)(nil)
	_ libdns.RecordSetter   = (*Provider)(nil)
)

// Provider is a wrapper of Client that implements RecordAppender, RecordDeleter,
// RecordGetter and RecordSetter interface.
type Provider struct {
	*Client
}

// GetRecords returns all the records in the DNS zone.
func (p *Provider) GetRecords(ctx context.Context, zone string) ([]libdns.Record, error) {
	domain := formatDomain(zone)
	req := NewQueryRecordsRequest(domain)
	resp, err := p.Client.QueryRecords(ctx, req)
	if err != nil {
		return nil, err
	}

	result := make([]libdns.Record, 0, len(resp.Records))
	for _, rec := range resp.Records {
		result = append(result, rec.libdnsRecord())
	}

	return result, nil
}

// AppendRecords creates the requested records in the given zone
// and returns the populated records that were created. It never
// changes existing records.
func (p *Provider) AppendRecords(ctx context.Context, zone string, recs []libdns.Record) ([]libdns.Record, error) {
	domain := formatDomain(zone)
	var result []libdns.Record
	for _, rec := range recs {
		subDomain := formatSubdomain(domain, rec.Name)
		req := NewCreateRecordRequest(domain, subDomain, rec.Type, rec.Value)
		req.TTL = formatTTL(rec.TTL)
		resp := &CreateRecordResponse{}
		err := p.Client.Do(ctx, req, resp)
		if err != nil {
			return nil, err
		}
		result = append(result, resp.Record.libdnsRecord())
	}
	return result, nil
}

// SetRecords updates the zone so that the records described in the
// input are reflected in the output. It may create or overwrite
// records or -- depending on the record type -- delete records to
// maintain parity with the input. No other records are affected.
// It returns the records which were set.
func (p *Provider) SetRecords(ctx context.Context, zone string, recs []libdns.Record) ([]libdns.Record, error) {
	domain := formatDomain(zone)
	var results []libdns.Record
	for _, rec := range recs {
		subdomain := formatSubdomain(domain, rec.Name)
		ttl := formatTTL(rec.TTL)
		if rec.ID == "" {
			queryReq := NewQueryRecordsRequest(domain)
			queryReq.Subdomain = subdomain
			queryReq.RecordType = rec.Type
			queryResp, err := p.Client.QueryRecords(ctx, queryReq)
			if err != nil {
				return nil, err
			}
			if len(queryResp.Records) == 0 {
				// record doesn't exist, create it.
				createReq := NewCreateRecordRequest(domain, subdomain, rec.Type, rec.Value)
				createReq.TTL = ttl
				createResp, err := p.Client.CreateRecord(ctx, createReq)
				if err != nil {
					return nil, err
				}
				rec.ID = createResp.Record.ID
				results = append(results, rec)
				continue
			}
			if len(queryResp.Records) > 1 {
				return nil, fmt.Errorf("unexpectedly found more than 1 record for %v", rec)
			}
			rec.ID = queryResp.Records[0].ID
		}
		// record exists, update it.
		updateReq := NewUpdateRecordRequest(domain, rec.ID, rec.Type, rec.Value)
		updateReq.Subdomain = subdomain
		updateReq.TTL = ttl
		updateResp, err := p.Client.UpdateRecord(ctx, updateReq)
		if err != nil {
			return nil, err
		}
		rec.ID = strconv.Itoa(updateResp.Record.ID)
		results = append(results, rec)
	}
	return results, nil
}

func (p *Provider) findRecordID(ctx context.Context, domain string, rec libdns.Record) (string, error) {
	queryReq := NewQueryRecordsRequest(domain)
	queryReq.Subdomain = formatSubdomain(domain, rec.Name)
	queryReq.RecordType = rec.Type
	queryResp, err := p.Client.QueryRecords(ctx, queryReq)
	if err != nil {
		return "", err
	}

	if len(queryResp.Records) == 0 {
		return "", fmt.Errorf("no such record %v", rec)
	}

	if len(queryResp.Records) > 1 {
		return "", fmt.Errorf("unexpectedly found more than 1 record for %v", rec)
	}

	return queryResp.Records[0].ID, nil
}

// DeleteRecords deletes the given records from the zone if they exist.
// It returns the records that were deleted.
func (p *Provider) DeleteRecords(ctx context.Context, zone string, recs []libdns.Record) ([]libdns.Record, error) {
	domain := formatDomain(zone)
	var result []libdns.Record
	for _, rec := range recs {
		if rec.ID == "" {
			id, err := p.findRecordID(ctx, domain, rec)
			if err != nil {
				return nil, err
			}
			rec.ID = id
		}
		req := NewDeleteRecordRequest(domain, rec.ID)
		_, err := p.Client.DeleteRecord(ctx, req)
		if err != nil {
			return nil, err
		}
		result = append(result, rec)
	}
	return result, nil
}
