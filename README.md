# Elastic Apm Operator
![Github Actions Workflow](https://github.com/unravellingtechnologies/elastic-apm-operator/actions/workflows/go.yml/badge.svg) [![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=unravellingtechnologies_zert&metric=alert_status)](https://sonarcloud.io/dashboard?id=unravellingtechnologies_elastic-apm-operator)

Kubernetes operator that automatically configures Elastic APM according to the configuration.

# Motivation
We use Elastic APM in a client as our APM tool of choice. Because we wanted a non-intrusive way (no code changes) to load the agent, we chose to use an initialisation container to load the agent's jar and then make use of the `-javaagent` flag. This works great, but we faced some challenges or detected possible improvements to our workflow:
* virtually all the information we pass on to the agent can be inferred from the deployment manifest or defaulted
* as the deployments got more and more automated (in our case using [FluxCD](https://fluxcd.io/)), the updating of the version, for example, became more error-prone
