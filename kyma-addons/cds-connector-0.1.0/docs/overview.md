---
title: Overview
type: Overview
---

## Training in the Now

SAP CDS is used for building user profiles using behavioural data like click streams.
## API Usage

Use the **SAP CDS Connector** to connect to SAP CDS system and be able to make API calls to the SAP CDS system from a lambda or a microservice. A sample URL to perform CRUD on CDS Profiles would be `process.env[{PREFIX}-GATEWAY_URL] + "/profiles"`.
Note that `PREFIX` is optional and used when specified while creating a service binding.
