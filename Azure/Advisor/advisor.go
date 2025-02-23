package Advisor

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor"
)

type Recommendation struct {
	RecName           string
	RecID             string
	ShortDescription  string
	ShortSolution     string
	Description       string
	Actions           []map[string]any
	ActionsRepo       ActionRepo
	Category          armadvisor.Category
	Impact            armadvisor.Impact
	ImpactedField     string
	ImpactedValue     string
	PotentialBenefits string
}

type ActionRepo struct {
	Description string
	ActionType  string
	Caption     string
	Link        string
	Metadata    ActionMetadata
}

type ActionMetadata struct {
	ID string
}

func GetAllAdvisories(recommendationsclientvariable *armadvisor.RecommendationsClient) (recommendationoutput map[string]Recommendation) {
	currentctx := context.Background()
	recommendationsclientvariable.Generate(currentctx, nil)
	currentpager := recommendationsclientvariable.NewListPager(nil)
	recommendationoutput = make(map[string]Recommendation)
	natext := "N/A"
	for currentpager.More() {
		currentpage, err := currentpager.NextPage(currentctx)
		if err != nil {
			fmt.Println("Unable to retrieve recommendations. Check authentication.", err)
		}
		for _, rec := range currentpage.Value {
			if rec.Name != nil && rec.ID != nil {
				var currentrec Recommendation
				currentrec.RecName = *rec.Name
				currentrec.RecID = *rec.ID
				if rec.Properties.ShortDescription.Problem != nil {
					currentrec.ShortDescription = *rec.Properties.ShortDescription.Problem
				} else {
					currentrec.ShortDescription = natext
				}
				if rec.Properties.ShortDescription.Solution != nil {
					currentrec.ShortSolution = *rec.Properties.ShortDescription.Solution
				} else {
					currentrec.ShortSolution = natext
				}
				if rec.Properties.Description != nil {
					currentrec.Description = *rec.Properties.Description
				} else {
					currentrec.Description = natext
				}
				if rec.Properties.Actions != nil {
					currentrec.Actions = *&rec.Properties.Actions
					for i := 0; i < len(currentrec.Actions); i++ {
						currentrec.ActionsRepo.Description = currentrec.Actions[i]["description"].(string)
						currentrec.ActionsRepo.ActionType = currentrec.Actions[i]["actionType"].(string)
						currentrec.ActionsRepo.Caption = currentrec.Actions[i]["caption"].(string)
						currentrec.ActionsRepo.Link = currentrec.Actions[i]["link"].(string)
						currentrec.ActionsRepo.Metadata.ID = currentrec.Actions[i]["metadata"].(string)
					}
				}
				if rec.Properties.Category != nil {
					currentrec.Category = *rec.Properties.Category
				}
				if rec.Properties.Impact != nil {
					currentrec.Impact = *rec.Properties.Impact
				}
				if rec.Properties.ImpactedField != nil {
					currentrec.ImpactedField = *rec.Properties.ImpactedField
				} else {
					currentrec.ImpactedField = natext
				}
				if rec.Properties.PotentialBenefits != nil {
					currentrec.PotentialBenefits = *rec.Properties.PotentialBenefits
				} else {
					currentrec.PotentialBenefits = natext
				}
				if rec.Properties.ImpactedValue != nil {
					currentrec.ImpactedValue = *rec.Properties.ImpactedValue
				} else {
					currentrec.ImpactedValue = natext
				}
				recommendationoutput[*rec.Name] = currentrec
			}
		}
	}
	return recommendationoutput
}
