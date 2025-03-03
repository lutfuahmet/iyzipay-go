package Iyzipay

import (
	"encoding/json"
	"github.com/lutfuahmet/iyzi/Request"
	"github.com/lutfuahmet/iyzi/Response"
)

func (i Iyzipay) InstallmentCheck(obj *Request.InstallmentRequest) (Response.InstallmentResponse, error) {

	result := Response.InstallmentResponse{}
	if resp, err := i.Client.R().
		SetHeaders(i.GetHeaders(*obj)).
		SetBody(obj).
		Post(i.GetURI("/payment/iyzipos/installment")); err == nil {
		err = json.Unmarshal(resp.Body(), &result)
		if err != nil {
			return result, err
		}
		return result, nil
	} else {
		return result, err
	}
}
