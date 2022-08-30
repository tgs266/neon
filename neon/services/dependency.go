package services

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/gnboorse/centipede"
	"github.com/tgs266/neon/neon/store/entities"
)

func buildConstraintString(dep entities.Dependency) string {
	out := []string{}
	if dep.MinVersion != "" {
		out = append(out, ">= "+dep.MinVersion)
	}
	if dep.MaxVersion != "" {
		out = append(out, "<= "+dep.MaxVersion)
	}
	return strings.Join(out, " ")
}

func resolveDependencies(products []entities.Product) map[string]*entities.Release {
	vars := centipede.Variables[*entities.Release]{}
	constraints := centipede.Constraints[*entities.Release]{}
	for _, p := range products {
		for _, r := range p.Releases {
			for _, d := range r.Dependencies {
				constr, _ := semver.NewConstraint(buildConstraintString(d))
				constraints = append(constraints, centipede.Constraint[*entities.Release]{
					Vars: centipede.VariableNames{centipede.VariableName(p.Name), centipede.VariableName(d.ProductName)},
					ConstraintFunction: func(variables *centipede.Variables[*entities.Release]) bool {
						if variables.Find(centipede.VariableName(p.Name)).Empty || variables.Find(centipede.VariableName(d.ProductName)).Empty {
							return true
						}

						value := variables.Find(centipede.VariableName(d.ProductName)).Value
						ver, _ := semver.NewVersion(value.ProductVersion)
						return constr.Check(ver)

					},
				})
			}
		}
		productVar := centipede.NewVariable(centipede.VariableName(p.Name), p.Releases)
		vars = append(vars, productVar)
	}

	solver := centipede.NewBackTrackingCSPSolver(vars, constraints)
	begin := time.Now()
	success, e := solver.Solve(context.TODO()) // run the solution
	elapsed := time.Since(begin)
	if e != nil {
		return nil
	}
	if success {
		fmt.Printf("Found solution in %s\n", elapsed)
		out := map[string]*entities.Release{}
		for _, variable := range solver.State.Vars {
			out[string(variable.Name)] = variable.Value
		}
		return out
	} else {
		return nil
	}
}
