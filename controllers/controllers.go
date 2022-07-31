package controllers

import "github.com/VictorAvelar/mollie-api-go/mollie"

func Create(client *mollie.Client, p mollie.Payment) (mollie.Payment, error) {
	payment, err := client.Payments.Create(p)
	if err != nil {
		return mollie.Payment{}, err
	}

	return payment, nil
}

func GetAll(client *mollie.Client, options *mollie.ListPaymentOptions) (mollie.PaymentList, error) {
	payments, err := client.Payments.List(options)
	if err != nil {
		return mollie.PaymentList{}, err
	}

	return payments, nil
}

func GetOne(client *mollie.Client, id string, options *mollie.PaymentOptions) (mollie.Payment, error) {
	payment, err := client.Payments.Get(id, options)
	if err != nil {
		return mollie.Payment{}, err
	}

	return payment, nil
}

func Update(client *mollie.Client, id string, p mollie.Payment) (mollie.Payment, error) {
	payment, err := client.Payments.Update(id, p)
	if err != nil {
		return mollie.Payment{}, err
	}

	return payment, nil
}

func Delete(client *mollie.Client, id string) (mollie.Payment, error) {
	payment, err := client.Payments.Cancel(id)
	if err != nil {
		return mollie.Payment{}, err
	}

	return payment, nil
}
