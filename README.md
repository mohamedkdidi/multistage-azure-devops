# Multistage Azure DevOps Pipeline

Simple build and release pipeline that builds a container, pushes it to ACR, deploys it to Azure Kubernetes Service linked to an environment.


This example based on an application Go (HTTP server) with multi-step Dockerfile which defines compiling the executable and then building the container to run.

Make sure you set your service connections for Azure Container Registry and Azure Kubernetes Service in azure-pipelines.yml file using this variables:

'''
  acrServiceConnexion: '' # Azure Container Registry service connexion
  aksServiceConnexion: '' #  Azure Kubernetes Service service connexion
  containerRegistry:   '' # your container registry name xxx.azurecr.io
'''