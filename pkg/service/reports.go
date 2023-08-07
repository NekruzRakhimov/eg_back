package service

import (
	"eg_back/models"
	"eg_back/pkg/repository"
	"fmt"
)

func GetReports(filter models.ReportFilter) (r models.Report, err error) {
	r.Active, err = repository.GetReportByStatus("active", filter)
	if err != nil {
		return models.Report{}, err
	}

	r.Liquidated, err = repository.GetReportByStatus("liquidated", filter)
	if err != nil {
		return models.Report{}, err
	}

	r.Stopped, err = repository.GetReportByStatus("stopped", filter)
	if err != nil {
		return models.Report{}, err
	}

	r.OnBankruptcy, err = repository.GetReportByStatus("on_bankruptcy", filter)
	if err != nil {
		return models.Report{}, err
	}

	r.LeasedOut, err = repository.GetReportByStatus("leased_out", filter)
	if err != nil {
		return models.Report{}, err
	}

	r.ReOrganized, err = repository.GetReportByStatus("re_organized", filter)
	if err != nil {
		return models.Report{}, err
	}

	r.CustomsDebtors, err = repository.GetReportByStatus("customs_debtors", filter)
	if err != nil {
		return models.Report{}, err
	}

	r.TaxDebtors, err = repository.GetReportByStatus("tax_debtors", filter)
	if err != nil {
		return models.Report{}, err
	}

	fmt.Println(">>>>", r)
	return r, nil
}
