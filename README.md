[![License: Apache 2](https://img.shields.io/badge/License-apache2-green.svg)](LICENSE)
[![operator](https://github.com/cernide/operator/actions/workflows/tests.yml/badge.svg)](https://github.com/cernide/operator/actions/workflows/tests.yml)
[![Slack](https://img.shields.io/badge/chat-on%20slack-aadada.svg?logo=slack&longCache=true)](https://polyaxon.com/slack/)
[![Docs](https://img.shields.io/badge/docs-stable-brightgreen.svg?style=flat)](https://polyaxon.com/docs/)
[![GitHub](https://img.shields.io/badge/issue_tracker-github-blue?logo=github)](https://github.com/cernide/cernide/issues)
[![GitHub](https://img.shields.io/badge/roadmap-github-blue?logo=github)](https://github.com/cernide/cernide/milestones)

<br>
<p align="center">
  <p align="center">
    <img src="https://raw.githubusercontent.com/polyaxon/polyaxon/master/artifacts/packages/operator.svg" alt="operator" height="100">
  </p>
</p>
<br>

# Machine Learning Operator & Controller for Kubernetes

## Introduction

Kubernetes offers the facility of extending it's API through the concept of 'Operators' ([Introducing Operators: Putting Operational Knowledge into Software](https://coreos.com/blog/introducing-operators.html)). This repository contains the resources and code to deploy an Polyaxon native CRDs using a native Operator for Kubernetes.

This project is a Kubernetes controller that manages and watches Customer Resource Definitions (CRDs) that define primitives to handle, operate and reconcile operations like: builds, jobs, experiments, distributed training, notebooks, tensorboards, kubeflow integrations, ...

![Operator Architecture](./artifacts/Operator-architecture.png)

## Kubeflow, Ray, and Dask operators

This Operator extends natively [Kubeflow-Operators](https://github.com/polyaxon/training-operator) (TFJob/PytorchJob/MXNet/XGBoost/MPI/Paddle), Dask Operator, Ray Operator.
