```bash
az role definition list \
    --query "[].{name:name, roleType:roleType, roleName:roleName}" \
    --output tsv > ~/lunch-and-learn-kube/AKS/az.role.definition.output.md
# big list of things!    
a7b1b19a-0e83-4fe5-935c-faaefbfd18c3	CustomRole	Avere Cluster Create
e078ab98-ef3a-4c9a-aba7-12f5172b45d0	CustomRole	Avere Cluster Runtime Operator
21d96096-b162-414a-8302-d8354f9d91b2	CustomRole	Azure Service Deploy Release Management Contributor
7b266cd7-0bba-4ae2-8423-90ede5e1e898	CustomRole	CAL-Custom-Role
b91f4c0b-46e3-47bb-a242-eecfe23b3b5b	CustomRole	Dsms Role (deprecated)
7aff565e-6c55-448d-83db-ccf482c6da2f	CustomRole	Dsms Role (do not use)
a48d7896-14b4-4889-afef-fbb65a96e5a2	CustomRole	ExpressRoute Administrator
9f15f5f5-77bd-413a-aa88-4b9c68b1e7bc	CustomRole	GenevaWarmPathResourceContributor
a48d7796-14b4-4889-afef-fbb65a93e5a2	CustomRole	masterreader
fd1bb084-1503-4bd2-99c0-630220046786	CustomRole	Microsoft OneAsset Reader
7fd64851-3279-459b-b614-e2b2ba760f5b	CustomRole	Office DevOps
a16c43ca-2d67-4dcd-9ded-6412f5edc51a	CustomRole	GenevaWarmPathStorageBlobContributor
94ddc4bc-25f5-4f3e-b527-c587da93cfe4	CustomRole	Azure Service Deploy Release Management Restricted Owner
# Key Vault Ref
87d31636-ad85-4caa-802d-1535972b5612	CustomRole	Azure Service Deploy Test Release Management Key Vault Secrets User
# Key Vault Ref
260691e6-68c2-47cf-bd4a-97d5fd4dbcd5	CustomRole	Azure Service Deploy Release Management Key Vault Secrets User
9b0f576e-fc2e-4256-9aa3-6fede171d599	CustomRole	AccessMonitoringReader
ebc6a08c-9436-41c0-98ea-9a98ce69e91c	CustomRole	CnAI Management Group Orchestrator Prod
fd6e57ea-fe3c-4f21-bd1e-de170a9a4971	CustomRole	AzSecPackUAPolicyResourceContributorCorpProd
47d81b3f-7139-46fc-a2e5-228265cc216c	CustomRole	SLNM-Admin-PolicyManagement-Prod
cbdd6655-6341-491c-b01b-8b202175cee4	CustomRole	SLNM-Teams-PolicyManagement-Prod
402b3d0b-63a7-4ca1-9ca1-56884c0dce6e	CustomRole	NetIso-Reserve-Role-Prod
6d224f7c-5c93-4466-b116-8c0c62090d0d	CustomRole	Policy Enforcement Role
8311e382-0749-4cb8-b61a-304f252e45ec	BuiltInRole	AcrPush
312a565d-c81f-4fd8-895a-4e21e48d571c	BuiltInRole	API Management Service Contributor
7f951dda-4ed3-4680-a7ca-43fe172d538d	BuiltInRole	AcrPull
6cef56e8-d556-48e5-a04f-b8e64114680f	BuiltInRole	AcrImageSigner
c2f4ef07-c644-48eb-af81-4b1b4947fb11	BuiltInRole	AcrDelete
cdda3590-29a3-44f6-95f2-9f980659eb04	BuiltInRole	AcrQuarantineReader
c8d4ff99-41c3-41a8-9f60-21dfdad59608	BuiltInRole	AcrQuarantineWriter
e022efe7-f5ba-4159-bbe4-b44f577e9b61	BuiltInRole	API Management Service Operator Role
71522526-b88f-4d52-b57f-d31fc3546d0d	BuiltInRole	API Management Service Reader Role
ae349356-3a1b-4a5e-921d-050484c6347e	BuiltInRole	Application Insights Component Contributor
08954f03-6346-4c2e-81c0-ec3a5cfae23b	BuiltInRole	Application Insights Snapshot Debugger
fd1bd22b-8476-40bc-a0bc-69b95687b9f3	BuiltInRole	Attestation Reader
4fe576fe-1146-4730-92eb-48519fa6bf9f	BuiltInRole	Automation Job Operator
5fb5aef8-1081-4b8e-bb16-9d5d0385bab5	BuiltInRole	Automation Runbook Operator
d3881f73-407a-4167-8283-e981cbba0404	BuiltInRole	Automation Operator
4f8fab4f-1852-4a58-a46a-8eaf358af14a	BuiltInRole	Avere Contributor
c025889f-8102-4ebf-b32c-fc0c6f0c6bd9	BuiltInRole	Avere Operator
0ab0b1a8-8aac-4efd-b8c2-3ee1fb270be8	BuiltInRole	Azure Kubernetes Service Cluster Admin Role
4abbcc35-e782-43d8-92c5-2d3f1bd2253f	BuiltInRole	Azure Kubernetes Service Cluster User Role
423170ca-a8f6-4b0f-8487-9e4eb8f49bfa	BuiltInRole	Azure Maps Data Reader
6f12a6df-dd06-4f3e-bcb1-ce8be600526a	BuiltInRole	Azure Stack Registration Owner
5e467623-bb1f-42f4-a55d-6e525e11384b	BuiltInRole	Backup Contributor
fa23ad8b-c56e-40d8-ac0c-ce449e1d2c64	BuiltInRole	Billing Reader
a795c7a0-d4a2-40c1-ae25-d81f01202912	BuiltInRole	Backup Reader
31a002a1-acaf-453e-8a5b-297c9ca1ea24	BuiltInRole	Blockchain Member Node Access (Preview)
5e3c6656-6cfa-4708-81fe-0de47ac73342	BuiltInRole	BizTalk Contributor
426e0c7f-0c7e-4658-b36f-ff54d6c29b45	BuiltInRole	CDN Endpoint Contributor
ec156ff8-a8d1-4d15-830c-5b80698ca432	BuiltInRole	CDN Profile Contributor
8f96442b-4075-438f-813d-ad51ab4019af	BuiltInRole	CDN Profile Reader
b34d265f-36f7-4a0d-a4d4-e158ca92e90f	BuiltInRole	Classic Network Contributor
86e8f5dc-a6e9-4c67-9d15-de283e8eac25	BuiltInRole	Classic Storage Account Contributor
985d6b00-f706-48f5-a6fe-d0ca12fb668d	BuiltInRole	Classic Storage Account Key Operator Service Role
9106cda0-8a86-4e81-b686-29a22c54effe	BuiltInRole	ClearDB MySQL DB Contributor
d73bb868-a0df-4d4d-bd69-98a00b01fccb	BuiltInRole	Classic Virtual Machine Contributor
a97b65f3-24c7-4388-baec-2e87135dc908	BuiltInRole	Cognitive Services User
b59867f0-fa02-499b-be73-45a86b5b3e1c	BuiltInRole	Cognitive Services Data Reader
25fbc0a9-bd7c-42a3-aa1a-3b75d497ee68	BuiltInRole	Cognitive Services Contributor
db7b14f2-5adf-42da-9f96-f2ee17bab5cb	BuiltInRole	CosmosBackupOperator
b24988ac-6180-42a0-ab88-20f7382dd24c	BuiltInRole	Contributor
fbdf93bf-df7d-467e-a4d2-9458aa1360c8	BuiltInRole	Cosmos DB Account Reader Role
434105ed-43f6-45c7-a02f-909b2ba83430	BuiltInRole	Cost Management Contributor
72fafb9e-0641-4937-9268-a91bfd8191a3	BuiltInRole	Cost Management Reader
add466c9-e687-43fc-8d98-dfcf8d720be5	BuiltInRole	Data Box Contributor
028f4ed7-e2a9-465e-a8f4-9c0ffdfdc027	BuiltInRole	Data Box Reader
673868aa-7521-48a0-acc6-0f60742d39f5	BuiltInRole	Data Factory Contributor
150f5e0c-0603-4f03-8c7f-cf70034c4e90	BuiltInRole	Data Purger
47b7735b-770e-4598-a7da-8b91488b4c88	BuiltInRole	Data Lake Analytics Developer
76283e04-6283-4c54-8f91-bcf1374a3c64	BuiltInRole	DevTest Labs User
5bd9cd88-fe45-4216-938b-f97437e15450	BuiltInRole	DocumentDB Account Contributor
befefa01-2a29-4197-83a8-272ff33ce314	BuiltInRole	DNS Zone Contributor
428e0ff0-5e57-4d9c-a221-2c70d0e0a443	BuiltInRole	EventGrid EventSubscription Contributor
2414bbcf-6497-4faf-8c65-045460748405	BuiltInRole	EventGrid EventSubscription Reader
b60367af-1334-4454-b71e-769d9a4f83d9	BuiltInRole	Graph Owner
8d8d5a11-05d3-4bda-a417-a08778121c7c	BuiltInRole	HDInsight Domain Services Contributor
03a6d094-3444-4b3d-88af-7477090a9e5e	BuiltInRole	Intelligent Systems Account Contributor
# Key Vault ref
f25e0fa2-a7c8-4377-a976-54943a77a395	BuiltInRole	Key Vault Contributor
ee361c5d-f7b5-4119-b4b6-892157c8f64c	BuiltInRole	Knowledge Consumer
b97fb8bc-a8b2-4522-a38b-dd33c7e65ead	BuiltInRole	Lab Creator
73c42c96-874c-492b-b04d-ab87d138a893	BuiltInRole	Log Analytics Reader
92aaf0da-9dab-42b6-94a3-d43ce8d16293	BuiltInRole	Log Analytics Contributor
515c2055-d9d4-4321-b1b9-bd0c9a0f79fe	BuiltInRole	Logic App Operator
87a39d53-fc1b-424a-814c-f7e04687dc9e	BuiltInRole	Logic App Contributor
c7393b34-138c-406f-901b-d8cf2b17e6ae	BuiltInRole	Managed Application Operator Role
b9331d33-8a36-4f8c-b097-4f54124fdb44	BuiltInRole	Managed Applications Reader
f1a07417-d97a-45cb-824c-7a7467783830	BuiltInRole	Managed Identity Operator
e40ec5ca-96e0-45a2-b4ff-59039f2c2b59	BuiltInRole	Managed Identity Contributor
5d58bcaf-24a5-4b20-bdb6-eed9f69fbe4c	BuiltInRole	Management Group Contributor
ac63b705-f282-497d-ac71-919bf39d939d	BuiltInRole	Management Group Reader
43d0d8ad-25c7-4714-9337-8ba259a9fe05	BuiltInRole	Monitoring Reader
4d97b98b-1d4f-4787-a291-c67834d212e7	BuiltInRole	Network Contributor
5d28c62d-5b37-4476-8438-e587778df237	BuiltInRole	New Relic APM Account Contributor
8e3af657-a8ff-443c-a75c-2fe8c4bcb635	BuiltInRole	Owner
acdd72a7-3385-48ef-bd42-f606fba81ae7	BuiltInRole	Reader
e0f68234-74aa-48ed-b826-c38b57376e17	BuiltInRole	Redis Cache Contributor
c12c1c16-33a1-487b-954d-41c89c60f349	BuiltInRole	Reader and Data Access
36243c78-bf99-498c-9df9-86d9f8d28608	BuiltInRole	Resource Policy Contributor
188a0f2f-5c9e-469b-ae67-2aa5ce574b94	BuiltInRole	Scheduler Job Collections Contributor
7ca78c08-252a-4471-8644-bb5ff32d4ba0	BuiltInRole	Search Service Contributor
e3d13bf0-dd5a-482e-ba6b-9b8433878d10	BuiltInRole	Security Manager (Legacy)
39bc4728-0917-49c7-9d2c-d95423bc2eb4	BuiltInRole	Security Reader
8bbe83f1-e2a6-4df7-8cb4-4e04d4e5c827	BuiltInRole	Spatial Anchors Account Contributor
6670b86e-a3f7-4917-ac9b-5d6ab1be4567	BuiltInRole	Site Recovery Contributor
494ae006-db33-4328-bf46-533a6560a3ca	BuiltInRole	Site Recovery Operator
5d51204f-eb77-4b1c-b86a-2ec626c49413	BuiltInRole	Spatial Anchors Account Reader
dbaa88c4-0c30-4179-9fb3-46319faa6149	BuiltInRole	Site Recovery Reader
70bbe301-9835-447d-afdd-19eb3167307c	BuiltInRole	Spatial Anchors Account Owner
4939a1f6-9ae0-4e48-a1e0-f2cbe897382d	BuiltInRole	SQL Managed Instance Contributor
9b7fa17d-e63e-47b0-bb0a-15c516ac86ec	BuiltInRole	SQL DB Contributor
056cd41c-7e88-42e1-933e-88ba6a50c9c3	BuiltInRole	SQL Security Manager
17d1049b-9a84-46fb-8f53-869881c3d3ab	BuiltInRole	Storage Account Contributor
6d8ee4ec-f05a-4a1d-8b00-a9b17e38b437	BuiltInRole	SQL Server Contributor
81a9662b-bebf-436f-a333-f67b29880f12	BuiltInRole	Storage Account Key Operator Service Role
ba92f5b4-2d11-453d-a403-e96b0029c9fe	BuiltInRole	Storage Blob Data Contributor
b7e6dc6d-f1e8-4753-8033-0f276bb0955b	BuiltInRole	Storage Blob Data Owner
2a2b9908-6ea1-4ae2-8e65-a410df84e7d1	BuiltInRole	Storage Blob Data Reader
974c5e8b-45b9-4653-ba55-5f855dd0fb88	BuiltInRole	Storage Queue Data Contributor
8a0f0c08-91a1-4084-bc3d-661d67233fed	BuiltInRole	Storage Queue Data Message Processor
c6a89b2d-59bc-44d0-9896-0f6e12d7b80a	BuiltInRole	Storage Queue Data Message Sender
19e7f393-937e-4f77-808e-94535e297925	BuiltInRole	Storage Queue Data Reader
cfd33db0-3dd1-45e3-aa9d-cdbdf3b6f24e	BuiltInRole	Support Request Contributor
a4b10055-b0c7-44c2-b00f-c7b5b3550cf7	BuiltInRole	Traffic Manager Contributor
18d7d88d-d35e-4fb5-a5c3-7773c20a72d9	BuiltInRole	User Access Administrator
9980e02c-c2be-4d73-94e8-173b1dc7cf3c	BuiltInRole	Virtual Machine Contributor
2cc479cb-7b4d-49a8-b449-8c00fd0f0a4b	BuiltInRole	Web Plan Contributor
de139f84-1756-47ae-9be6-808fbbe84772	BuiltInRole	Website Contributor
090c5cfd-751d-490a-894a-3ce6f1109419	BuiltInRole	Azure Service Bus Data Owner
f526a384-b230-433a-b45c-95f59c4a2dec	BuiltInRole	Azure Event Hubs Data Owner
bbf86eb8-f7b4-4cce-96e4-18cddf81d86e	BuiltInRole	Attestation Contributor
61ed4efc-fab3-44fd-b111-e24485cc132a	BuiltInRole	HDInsight Cluster Operator
230815da-be43-4aae-9cb4-875f7bd000aa	BuiltInRole	Cosmos DB Operator
48b40c6e-82e0-4eb3-90d5-19e40f49b624	BuiltInRole	Hybrid Server Resource Administrator
5d1e5ee4-7c68-4a71-ac8b-0739630a3dfb	BuiltInRole	Hybrid Server Onboarding
a638d3c7-ab3a-418d-83e6-5f17a39d4fde	BuiltInRole	Azure Event Hubs Data Receiver
2b629674-e913-4c01-ae53-ef4638d8f975	BuiltInRole	Azure Event Hubs Data Sender
4f6d3b9b-027b-4f4c-9142-0e5a2a2247e0	BuiltInRole	Azure Service Bus Data Receiver
69a216fc-b8fb-44d8-bc22-1f3c2cd27a39	BuiltInRole	Azure Service Bus Data Sender
aba4ae5f-2193-4029-9191-0cb91df5e314	BuiltInRole	Storage File Data SMB Share Reader
0c867c2a-1d8c-454a-a3db-ab2ea1bdc8bb	BuiltInRole	Storage File Data SMB Share Contributor
b12aa53e-6015-4669-85d0-8515ebb3ae7f	BuiltInRole	Private DNS Zone Contributor
db58b8e5-c6ad-4a2a-8342-4190687cbf4a	BuiltInRole	Storage Blob Delegator
1d18fff3-a72a-46b5-b4a9-0b38a3cd7e63	BuiltInRole	Desktop Virtualization User
a7264617-510b-434b-a828-9731dc254ea7	BuiltInRole	Storage File Data SMB Share Elevated Contributor
41077137-e803-4205-871c-5a86e6a753b4	BuiltInRole	Blueprint Contributor
437d2ced-4a38-4302-8479-ed2bcb43d090	BuiltInRole	Blueprint Operator
ab8e14d6-4a74-4a29-9ba8-549422addade	BuiltInRole	Microsoft Sentinel Contributor
3e150937-b8fe-4cfb-8069-0eaf05ecd056	BuiltInRole	Microsoft Sentinel Responder
8d289c81-5878-46d4-8554-54e1e3d8b5cb	BuiltInRole	Microsoft Sentinel Reader
66bb4e9e-b016-4a94-8249-4c0511c2be84	BuiltInRole	Policy Insights Data Writer (Preview)
04165923-9d83-45d5-8227-78b77b0a687e	BuiltInRole	SignalR AccessKey Reader
8cf5e20a-e4b2-4e9d-b3a1-5ceb692c2761	BuiltInRole	SignalR/Web PubSub Contributor
b64e21ea-ac4e-4cdf-9dc9-5b892992bee7	BuiltInRole	Azure Connected Machine Onboarding
91c1777a-f3dc-4fae-b103-61d183457e46	BuiltInRole	Managed Services Registration assignment Delete Role
5ae67dd6-50cb-40e7-96ff-dc2bfa4b606b	BuiltInRole	App Configuration Data Owner
516239f1-63e1-4d78-a4de-a74fb236a071	BuiltInRole	App Configuration Data Reader
34e09817-6cbe-4d01-b1a2-e0eac5743d41	BuiltInRole	Kubernetes Cluster - Azure Arc Onboarding
7f646f1b-fa08-80eb-a22b-edd6ce5c915c	BuiltInRole	Experimentation Contributor
466ccd10-b268-4a11-b098-b4849f024126	BuiltInRole	Cognitive Services QnA Maker Reader
f4cc2bf9-21be-47a1-bdf1-5c5804381025	BuiltInRole	Cognitive Services QnA Maker Editor
7f646f1b-fa08-80eb-a33b-edd6ce5c915c	BuiltInRole	Experimentation Administrator
3df8b902-2a6f-47c7-8cc5-360e9b272a7e	BuiltInRole	Remote Rendering Administrator
d39065c4-c120-43c9-ab0a-63eed9795f0a	BuiltInRole	Remote Rendering Client
641177b8-a67a-45b9-a033-47bc880bb21e	BuiltInRole	Managed Application Contributor Role
612c2aa1-cb24-443b-ac28-3ab7272de6f5	BuiltInRole	Security Assessment Contributor
4a9ae827-6dc8-4573-8ac7-8239d42aa03f	BuiltInRole	Tag Contributor
c7aa55d3-1abb-444a-a5ca-5e51e485d6ec	BuiltInRole	Integration Service Environment Developer
a41e2c5b-bd99-4a07-88f4-9bf657a760b8	BuiltInRole	Integration Service Environment Contributor
ed7f3fbd-7b88-4dd4-9017-9adb7ce333f8	BuiltInRole	Azure Kubernetes Service Contributor Role
d57506d4-4c8d-48b1-8587-93c323f6a5a3	BuiltInRole	Azure Digital Twins Data Reader
bcd981a7-7f74-457b-83e1-cceb9e632ffe	BuiltInRole	Azure Digital Twins Data Owner
350f8d15-c687-4448-8ae1-157740a3936d	BuiltInRole	Hierarchy Settings Administrator
5a1fc7df-4bf1-4951-a576-89034ee01acd	BuiltInRole	FHIR Data Contributor
3db33094-8700-4567-8da5-1501d4e7e843	BuiltInRole	FHIR Data Exporter
4c8d0bbc-75d3-4935-991f-5f3c56d81508	BuiltInRole	FHIR Data Reader
3f88fce4-5892-4214-ae73-ba5294559913	BuiltInRole	FHIR Data Writer
49632ef5-d9ac-41f4-b8e7-bbe587fa74a1	BuiltInRole	Experimentation Reader
4dd61c23-6743-42fe-a388-d8bdd41cb745	BuiltInRole	Object Understanding Account Owner
8f5e0ce6-4f7b-4dcf-bddf-e6f48634a204	BuiltInRole	Azure Maps Data Contributor
c1ff6cc2-c111-46fe-8896-e0ef812ad9f3	BuiltInRole	Cognitive Services Custom Vision Contributor
5c4089e1-6d96-4d2f-b296-c1bc7137275f	BuiltInRole	Cognitive Services Custom Vision Deployment
88424f51-ebe7-446f-bc41-7fa16989e96c	BuiltInRole	Cognitive Services Custom Vision Labeler
93586559-c37d-4a6b-ba08-b9f0940c2d73	BuiltInRole	Cognitive Services Custom Vision Reader
0a5ae4ab-0d65-4eeb-be61-29fc9b54394b	BuiltInRole	Cognitive Services Custom Vision Trainer
# Key Vautl Refs, lots of options here
00482a5a-887f-4fb3-b363-3b7fe8e74483	BuiltInRole	Key Vault Administrator
12338af0-0e69-4776-bea7-57ae8d297424	BuiltInRole	Key Vault Crypto User
b86a8fe4-44ce-4948-aee5-eccb2c155cd7	BuiltInRole	Key Vault Secrets Officer
4633458b-17de-408a-b874-0445c86b69e6	BuiltInRole	Key Vault Secrets User
a4417e6f-fecd-4de8-b567-7b0420556985	BuiltInRole	Key Vault Certificates Officer
21090545-7ca7-4776-b22c-e363652d74d2	BuiltInRole	Key Vault Reader
e147488a-f6f5-4113-8e2d-b22465e65bf6	BuiltInRole	Key Vault Crypto Service Encryption User
63f0a09d-1495-4db4-a681-037d84835eb4	BuiltInRole	Azure Arc Kubernetes Viewer
5b999177-9696-4545-85c7-50de3797e5a1	BuiltInRole	Azure Arc Kubernetes Writer
8393591c-06b9-48a2-a542-1bd6b377f6a2	BuiltInRole	Azure Arc Kubernetes Cluster Admin
dffb1e0c-446f-4dde-a09f-99eb5cc68b96	BuiltInRole	Azure Arc Kubernetes Admin
b1ff04bb-8a4e-4dc4-8eb5-8693973ce19b	BuiltInRole	Azure Kubernetes Service RBAC Cluster Admin
3498e952-d568-435e-9b2c-8d77e338d7f7	BuiltInRole	Azure Kubernetes Service RBAC Admin
7f6c6a51-bcf8-42ba-9220-52d62157d7db	BuiltInRole	Azure Kubernetes Service RBAC Reader
a7ffa36f-339b-4b5c-8bdf-e2c188b2c0eb	BuiltInRole	Azure Kubernetes Service RBAC Writer
82200a5b-e217-47a5-b665-6d8765ee745b	BuiltInRole	Services Hub Operator
d18777c0-1514-4662-8490-608db7d334b6	BuiltInRole	Object Understanding Account Reader
fd53cd77-2268-407a-8f46-7e7863d0f521	BuiltInRole	SignalR REST API Owner
daa9e50b-21df-454c-94a6-a8050adab352	BuiltInRole	Collaborative Data Contributor
e9dba6fb-3d52-4cf0-bce3-f06ce71b9e0f	BuiltInRole	Device Update Reader
02ca0879-e8e4-47a5-a61e-5c618b76e64a	BuiltInRole	Device Update Administrator
0378884a-3af5-44ab-8323-f5b22f9f3c98	BuiltInRole	Device Update Content Administrator
d1ee9a80-8b14-47f0-bdc2-f4a351625a7b	BuiltInRole	Device Update Content Reader
cb43c632-a144-4ec5-977c-e80c4affc34a	BuiltInRole	Cognitive Services Metrics Advisor Administrator
3b20f47b-3825-43cb-8114-4bd2201156a8	BuiltInRole	Cognitive Services Metrics Advisor User
2c56ea50-c6b3-40a6-83c0-9d98858bc7d2	BuiltInRole	Schema Registry Reader (Preview)
5dffeca3-4936-4216-b2bc-10343a5abb25	BuiltInRole	Schema Registry Contributor (Preview)
7ec7ccdc-f61e-41fe-9aaf-980df0a44eba	BuiltInRole	AgFood Platform Service Reader
8508508a-4469-4e45-963b-2518ee0bb728	BuiltInRole	AgFood Platform Service Contributor
f8da80de-1ff9-4747-ad80-a19b7f6079e3	BuiltInRole	AgFood Platform Service Admin
18500a29-7fe2-46b2-a342-b16a415e101d	BuiltInRole	Managed HSM contributor
0b555d9b-b4a7-4f43-b330-627f0e5be8f0	BuiltInRole	Security Detonation Chamber Submitter
ddde6b66-c0df-4114-a159-3618637b3035	BuiltInRole	SignalR REST API Reader
7e4f1700-ea5a-4f59-8f37-079cfe29dce3	BuiltInRole	SignalR Service Owner
f7b75c60-3036-4b75-91c3-6b41c27c1689	BuiltInRole	Reservation Purchaser
635dd51f-9968-44d3-b7fb-6d9a6bd613ae	BuiltInRole	AzureML Metrics Writer (preview)
e5e2a7ff-d759-4cd2-bb51-3152d37e2eb1	BuiltInRole	Storage Account Backup Contributor
6188b7c9-7d01-4f99-a59f-c88b630326c0	BuiltInRole	Experimentation Metric Contributor
9ef4ef9c-a049-46b0-82ab-dd8ac094c889	BuiltInRole	Project Babylon Data Curator
c8d896ba-346d-4f50-bc1d-7d1c84130446	BuiltInRole	Project Babylon Data Reader
05b7651b-dc44-475e-b74d-df3db49fae0f	BuiltInRole	Project Babylon Data Source Administrator
ca6382a4-1721-4bcf-a114-ff0c70227b6b	BuiltInRole	Application Group Contributor
49a72310-ab8d-41df-bbb0-79b649203868	BuiltInRole	Desktop Virtualization Reader
082f0a83-3be5-4ba1-904c-961cca79b387	BuiltInRole	Desktop Virtualization Contributor
21efdde3-836f-432b-bf3d-3e8e734d4b2b	BuiltInRole	Desktop Virtualization Workspace Contributor
ea4bfff8-7fb4-485a-aadd-d4129a0ffaa6	BuiltInRole	Desktop Virtualization User Session Operator
2ad6aaab-ead9-4eaa-8ac5-da422f562408	BuiltInRole	Desktop Virtualization Session Host Operator
ceadfde2-b300-400a-ab7b-6143895aa822	BuiltInRole	Desktop Virtualization Host Pool Reader
e307426c-f9b6-4e81-87de-d99efb3c32bc	BuiltInRole	Desktop Virtualization Host Pool Contributor
aebf23d0-b568-4e86-b8f9-fe83a2c6ab55	BuiltInRole	Desktop Virtualization Application Group Reader
86240b0e-9422-4c43-887b-b61143f32ba8	BuiltInRole	Desktop Virtualization Application Group Contributor
0fa44ee9-7a7d-466b-9bb2-2bf446b1204d	BuiltInRole	Desktop Virtualization Workspace Reader
3e5e47e6-65f7-47ef-90b5-e5dd4d455f24	BuiltInRole	Disk Backup Reader
b50d9833-a0cb-478e-945f-707fcc997c13	BuiltInRole	Disk Restore Operator
7efff54f-a5b4-42b5-a1c5-5411624893ce	BuiltInRole	Disk Snapshot Contributor
5548b2cf-c94c-4228-90ba-30851930a12f	BuiltInRole	Microsoft.Kubernetes connected cluster role
a37b566d-3efa-4beb-a2f2-698963fa42ce	BuiltInRole	Security Detonation Chamber Submission Manager
352470b3-6a9c-4686-b503-35deb827e500	BuiltInRole	Security Detonation Chamber Publisher
7a6f0e70-c033-4fb1-828c-08514e5f4102	BuiltInRole	Collaborative Runtime Operator
5432c526-bc82-444a-b7ba-57c5b0b5b34f	BuiltInRole	CosmosRestoreOperator
a1705bd2-3a8f-45a5-8683-466fcfd5cc24	BuiltInRole	FHIR Data Converter
0e5f05e5-9ab9-446b-b98d-1e2157c94125	BuiltInRole	Quota Request Operator
1e241071-0855-49ea-94dc-649edcd759de	BuiltInRole	EventGrid Contributor
28241645-39f8-410b-ad48-87863e2951d5	BuiltInRole	Security Detonation Chamber Reader
4a167cdf-cb95-4554-9203-2347fe489bd9	BuiltInRole	Object Anchors Account Reader
ca0835dd-bacc-42dd-8ed2-ed5e7230d15b	BuiltInRole	Object Anchors Account Owner
d17ce0a2-0697-43bc-aac5-9113337ab61c	BuiltInRole	WorkloadBuilder Migration Agent Role
b5537268-8956-4941-a8f0-646150406f0c	BuiltInRole	Azure Spring Cloud Data Reader
0e75ca1e-0464-4b4d-8b93-68208a576181	BuiltInRole	Cognitive Services Speech Contributor
9894cab4-e18a-44aa-828b-cb588cd6f2d7	BuiltInRole	Cognitive Services Face Recognizer
054126f8-9a2b-4f1c-a9ad-eca461f08466	BuiltInRole	Media Services Account Administrator
532bc159-b25e-42c0-969e-a1d439f60d77	BuiltInRole	Media Services Live Events Administrator
e4395492-1534-4db2-bedf-88c14621589c	BuiltInRole	Media Services Media Operator
c4bba371-dacd-4a26-b320-7250bca963ae	BuiltInRole	Media Services Policy Administrator
99dba123-b5fe-44d5-874c-ced7199a5804	BuiltInRole	Media Services Streaming Endpoints Administrator
1ec5b3c1-b17e-4e25-8312-2acb3c3c5abf	BuiltInRole	Stream Analytics Query Tester
a2138dac-4907-4679-a376-736901ed8ad8	BuiltInRole	AnyBuild Builder
b447c946-2db7-41ec-983d-d8bf3b1c77e3	BuiltInRole	IoT Hub Data Reader
494bdba2-168f-4f31-a0a1-191d2f7c028c	BuiltInRole	IoT Hub Twin Contributor
4ea46cd5-c1b2-4a8e-910b-273211f9ce47	BuiltInRole	IoT Hub Registry Contributor
4fc6c259-987e-4a07-842e-c321cc9d413f	BuiltInRole	IoT Hub Data Contributor
15e0f5a1-3450-4248-8e25-e2afe88a9e85	BuiltInRole	Test Base Reader
1407120a-92aa-4202-b7e9-c0e197c71c8f	BuiltInRole	Search Index Data Reader
8ebe5a00-799e-43f5-93ac-243d3dce84a7	BuiltInRole	Search Index Data Contributor
76199698-9eea-4c19-bc75-cec21354c6b6	BuiltInRole	Storage Table Data Reader
0a9a7e1f-b9d0-4cc4-a60d-0319b160aaa3	BuiltInRole	Storage Table Data Contributor
e89c7a3c-2f64-4fa1-a847-3e4c9ba4283a	BuiltInRole	DICOM Data Reader
58a3b984-7adf-4c20-983a-32417c86fbc8	BuiltInRole	DICOM Data Owner
d5a91429-5739-47e2-a06b-3470a27159e7	BuiltInRole	EventGrid Data Sender
60fc6e62-5479-42d4-8bf4-67625fcc2840	BuiltInRole	Disk Pool Operator
f6c7c914-8db3-469d-8ca1-694a8f32e121	BuiltInRole	AzureML Data Scientist
22926164-76b3-42b3-bc55-97df8dab3e41	BuiltInRole	Grafana Admin
e8113dce-c529-4d33-91fa-e9b972617508	BuiltInRole	Azure Connected SQL Server Onboarding
26baccc8-eea7-41f1-98f4-1762cc7f685d	BuiltInRole	Azure Relay Sender
2787bf04-f1f5-4bfe-8383-c8a24483ee38	BuiltInRole	Azure Relay Owner
26e0b698-aa6d-4085-9386-aadae190014d	BuiltInRole	Azure Relay Listener
60921a7e-fef1-4a43-9b16-a26c52ad4769	BuiltInRole	Grafana Viewer
a79a5197-3a5c-4973-a920-486035ffd60f	BuiltInRole	Grafana Editor
f353d9bd-d4a6-484e-a77a-8050b599b867	BuiltInRole	Automation Contributor
85cb6faf-e071-4c9b-8136-154b5a04f717	BuiltInRole	Kubernetes Extension Contributor
10745317-c249-44a1-a5ce-3a4353c0bbd8	BuiltInRole	Device Provisioning Service Data Reader
dfce44e4-17b7-4bd1-a6d1-04996ec95633	BuiltInRole	Device Provisioning Service Data Contributor
2837e146-70d7-4cfd-ad55-7efa6464f958	BuiltInRole	Trusted Signing Certificate Profile Signer
cff1b556-2399-4e7e-856d-a8f754be7b65	BuiltInRole	Azure Spring Cloud Service Registry Reader
f5880b48-c26d-48be-b172-7927bfa1c8f1	BuiltInRole	Azure Spring Cloud Service Registry Contributor
d04c6db6-4947-4782-9e91-30a88feb7be7	BuiltInRole	Azure Spring Cloud Config Server Reader
a06f5c24-21a7-4e1a-aa2b-f19eb6684f5b	BuiltInRole	Azure Spring Cloud Config Server Contributor
6ae96244-5829-4925-a7d3-5975537d91dd	BuiltInRole	Azure VM Managed identities restore Contributor
6be48352-4f82-47c9-ad5e-0acacefdb005	BuiltInRole	Azure Maps Search and Render Data Reader
dba33070-676a-4fb0-87fa-064dc56ff7fb	BuiltInRole	Azure Maps Contributor
b748a06d-6150-4f8a-aaa9-ce3940cd96cb	BuiltInRole	Azure Arc VMware VM Contributor
ce551c02-7c42-47e0-9deb-e3b6fc3a9a83	BuiltInRole	Azure Arc VMware Private Cloud User
ddc140ed-e463-4246-9145-7c664192013f	BuiltInRole	Azure Arc VMware Administrator role 
f72c8140-2111-481c-87ff-72b910f6e3f8	BuiltInRole	Cognitive Services LUIS Owner
7628b7b8-a8b2-4cdc-b46f-e9b35248918e	BuiltInRole	Cognitive Services Language Reader
f2310ca1-dc64-4889-bb49-c8e0fa3d47a8	BuiltInRole	Cognitive Services Language Writer
f07febfe-79bc-46b1-8b37-790e26e6e498	BuiltInRole	Cognitive Services Language Owner
18e81cdc-4e98-4e29-a639-e7d10c5a6226	BuiltInRole	Cognitive Services LUIS Reader
6322a993-d5c9-4bed-b113-e49bbea25b27	BuiltInRole	Cognitive Services LUIS Writer
a9a19cc5-31f4-447c-901f-56c0bb18fcaf	BuiltInRole	PlayFab Reader
749a398d-560b-491b-bb21-08924219302e	BuiltInRole	Load Test Contributor
45bb0b16-2f0c-4e78-afaa-a07599b003f6	BuiltInRole	Load Test Owner
0c8b84dc-067c-4039-9615-fa1a4b77c726	BuiltInRole	PlayFab Contributor
3ae3fb29-0000-4ccd-bf80-542e7b26e081	BuiltInRole	Load Test Reader
b2de6794-95db-4659-8781-7e080d3f2b9d	BuiltInRole	Cognitive Services Immersive Reader User
f69b8690-cc87-41d6-b77a-a4bc3c0a966f	BuiltInRole	Lab Services Contributor
2a5c394f-5eb7-4d4f-9c8e-e8eae39faebc	BuiltInRole	Lab Services Reader
ce40b423-cede-4313-a93f-9b28290b72e1	BuiltInRole	Lab Assistant
a36e6959-b6be-4b12-8e9f-ef4b474d304d	BuiltInRole	Lab Operator
5daaa2af-1fe8-407c-9122-bba179798270	BuiltInRole	Lab Contributor
fb1c8493-542b-48eb-b624-b4c8fea62acd	BuiltInRole	Security Admin
12cf5a90-567b-43ae-8102-96cf46c7d9b4	BuiltInRole	Web PubSub Service Owner
bfb1c7d2-fb1a-466b-b2ba-aee63b92deaf	BuiltInRole	Web PubSub Service Reader
420fcaa2-552c-430f-98ca-3264be4806c7	BuiltInRole	SignalR App Server
fb879df8-f326-4884-b1cf-06f3ad86be52	BuiltInRole	Virtual Machine User Login
1c0163c0-47e6-4577-8991-ea5c82e286e4	BuiltInRole	Virtual Machine Administrator Login
cd570a14-e51a-42ad-bac8-bafd67325302	BuiltInRole	Azure Connected Machine Resource Administrator
00c29273-979b-4161-815c-10b084fb9324	BuiltInRole	Backup Operator
e8ddcd69-c73f-4f9f-9844-4100522f16ad	BuiltInRole	Workbook Contributor
b279062a-9be3-42a0-92ae-8b3cf002ec4d	BuiltInRole	Workbook Reader
749f88d5-cbae-40b8-bcfc-e573ddc772fa	BuiltInRole	Monitoring Contributor
3913510d-42f4-4e42-8a64-420c390055eb	BuiltInRole	Monitoring Metrics Publisher
8a3c2885-9b38-4fd2-9d99-91af537c1347	BuiltInRole	Purview role 1 (Deprecated)
200bba9e-f0c8-430f-892b-6f0794863803	BuiltInRole	Purview role 2 (Deprecated)
ff100721-1b9d-43d8-af52-42b69c1272db	BuiltInRole	Purview role 3 (Deprecated)
b8b15564-4fa6-4a59-ab12-03e1d9594795	BuiltInRole	Autonomous Development Platform Data Contributor (Preview)
27f8b550-c507-4db9-86f2-f4b8e816d59d	BuiltInRole	Autonomous Development Platform Data Owner (Preview)
d63b75f7-47ea-4f27-92ac-e0d173aaf093	BuiltInRole	Autonomous Development Platform Data Reader (Preview)
14b46e9e-c2b7-41b4-b07b-48a6ebf60603	BuiltInRole	Key Vault Crypto Officer
49e2f5d2-7741-4835-8efa-19e1fe35e47f	BuiltInRole	Device Update Deployments Reader
e4237640-0e3d-4a46-8fda-70bc94856432	BuiltInRole	Device Update Deployments Administrator
67d33e57-3129-45e6-bb0b-7cc522f762fa	BuiltInRole	Azure Arc VMware Private Clouds Onboarding
4e9b8407-af2e-495b-ae54-bb60a55b1b5a	BuiltInRole	Chamber Admin
f4c81013-99ee-4d62-a7ee-b3f1f648599a	BuiltInRole	Microsoft Sentinel Automation Contributor
871e35f6-b5c1-49cc-a043-bde969a0f2cd	BuiltInRole	CDN Endpoint Reader
4447db05-44ed-4da3-ae60-6cbece780e32	BuiltInRole	Chamber User
f2dc8367-1007-4938-bd23-fe263f013447	BuiltInRole	Cognitive Services Speech User
a6333a3e-0164-44c3-b281-7a577aff287f	BuiltInRole	Windows Admin Center Administrator Login
18ed5180-3e48-46fd-8541-4ea054d57064	BuiltInRole	Azure Kubernetes Service Policy Add-on Deployment
088ab73d-1256-47ae-bea9-9de8e7131f31	BuiltInRole	Guest Configuration Resource Contributor
361898ef-9ed1-48c2-849c-a832951106bb	BuiltInRole	Domain Services Reader
eeaeda52-9324-47f6-8069-5d5bade478b2	BuiltInRole	Domain Services Contributor
0f2ebee7-ffd4-4fc0-b3b7-664099fdad5d	BuiltInRole	DNS Resolver Contributor
00493d72-78f6-4148-b6c5-d3ce8e4799dd	BuiltInRole	Azure Arc Enabled Kubernetes Cluster User Role
959f8984-c045-4866-89c7-12bf9737be2e	BuiltInRole	Data Operator for Managed Disks
6b77f0a0-0d89-41cc-acd1-579c22c17a67	BuiltInRole	AgFood Platform Sensor Partner Contributor
1ef6a3be-d0ac-425d-8c01-acb62866290b	BuiltInRole	Compute Gallery Sharing Admin
cd08ab90-6b14-449c-ad9a-8f8e549482c6	BuiltInRole	Scheduled Patching Contributor
45d50f46-0b78-4001-a660-4198cbe8cd05	BuiltInRole	DevCenter Dev Box User
331c37c6-af14-46d9-b9f4-e1909e1b95a0	BuiltInRole	DevCenter Project Admin
602da2ba-a5c2-41da-b01d-5360126ab525	BuiltInRole	Virtual Machine Local User Login
6aac74c4-6311-40d2-bbdd-7d01e7c6e3a9	BuiltInRole	Azure Arc ScVmm Private Clouds Onboarding
e582369a-e17b-42a5-b10c-874c387c530b	BuiltInRole	Azure Arc ScVmm VM Contributor
c0781e91-8102-4553-8951-97c6d4243cda	BuiltInRole	Azure Arc ScVmm Private Cloud User
a92dfd61-77f9-4aec-a531-19858b406c87	BuiltInRole	Azure Arc ScVmm Administrator role
fd036e6b-1266-47a0-b0bb-a05d04831731	BuiltInRole	HDInsight on AKS Cluster Admin
7656b436-37d4-490a-a4ab-d39f838f0042	BuiltInRole	HDInsight on AKS Cluster Pool Admin
4465e953-8ced-4406-a58e-0f6e3f3b530b	BuiltInRole	FHIR Data Importer
bcf28286-af25-4c81-bb6f-351fcab5dbe9	BuiltInRole	HDInsight on AKS Cluster Operator
c031e6a8-4391-4de0-8d69-4706a7ed3729	BuiltInRole	API Management Developer Portal Content Editor
d24ecba3-c1f4-40fa-a7bb-4588a071e8fd	BuiltInRole	VM Scanner Operator
80dcbedb-47ef-405d-95bd-188a1b4ac406	BuiltInRole	Elastic SAN Owner
af6a70f8-3c9f-4105-acf1-d719e9fca4ca	BuiltInRole	Elastic SAN Reader
a959dbd1-f747-45e3-8ba6-dd80f235f97c	BuiltInRole	Desktop Virtualization Virtual Machine Contributor
40c5ff49-9181-41f8-ae61-143b0e78555e	BuiltInRole	Desktop Virtualization Power On Off Contributor
489581de-a3bd-480d-9518-53dea7416b33	BuiltInRole	Desktop Virtualization Power On Contributor
a8281131-f312-4f34-8d98-ae12be9f0d23	BuiltInRole	Elastic SAN Volume Group Owner
76cc9ee4-d5d3-4a45-a930-26add3d73475	BuiltInRole	Access Review Operator Service Role
4339b7cf-9826-4e41-b4ed-c7f4505dac08	BuiltInRole	Trusted Signing Identity Verifier
a2c4a527-7dc0-4ee3-897b-403ade70fafb	BuiltInRole	Video Indexer Restricted Viewer
b0d8363b-8ddd-447d-831f-62ca05bff136	BuiltInRole	Monitoring Data Reader
18ab4d3d-a1bf-4477-8ad9-8359bc988f69	BuiltInRole	Azure Kubernetes Fleet Manager RBAC Cluster Admin
30b27cfc-9c84-438e-b0ce-70e35255df80	BuiltInRole	Azure Kubernetes Fleet Manager RBAC Reader
434fb43a-c01c-447e-9f67-c3ad923cfaba	BuiltInRole	Azure Kubernetes Fleet Manager RBAC Admin
5af6afb3-c06c-4fa4-8848-71a8aee05683	BuiltInRole	Azure Kubernetes Fleet Manager RBAC Writer
63bb64ad-9799-4770-b5c3-24ed299a07bf	BuiltInRole	Azure Kubernetes Fleet Manager Contributor Role
ba79058c-0414-4a34-9e42-c3399d80cd5a	BuiltInRole	Kubernetes Namespace User
c6decf44-fd0a-444c-a844-d653c394e7ab	BuiltInRole	Data Labeling - Labeler
f58310d9-a9f6-439a-9e8d-f62e7b41a168	BuiltInRole	Role Based Access Control Administrator
1c9b6475-caf0-4164-b5a1-2142a7116f4b	BuiltInRole	Template Spec Contributor
392ae280-861d-42bd-9ea5-08ee6d83b80e	BuiltInRole	Template Spec Reader
51d6186e-6489-4900-b93f-92e23144cca5	BuiltInRole	Microsoft Sentinel Playbook Operator
18e40d4e-8d2e-438d-97e1-9528336e149c	BuiltInRole	Deployment Environments User
80558df3-64f9-4c0f-b32d-e5094b036b0b	BuiltInRole	Azure Spring Apps Connect Role
a99b0159-1064-4c22-a57b-c9b3caa1c054	BuiltInRole	Azure Spring Apps Remote Debugging Role
e503ece1-11d0-4e8e-8e2c-7a6c3bf38815	BuiltInRole	AzureML Compute Operator
1823dd4f-9b8c-4ab6-ab4e-7397a3684615	BuiltInRole	AzureML Registry User
aabbc5dd-1af0-458b-a942-81af88f9c138	BuiltInRole	Azure Center for SAP solutions service role
05352d14-a920-4328-a0de-4cbe7430e26b	BuiltInRole	Azure Center for SAP solutions reader
7b0c7e81-271f-4c71-90bf-e30bdfdbc2f7	BuiltInRole	Azure Center for SAP solutions administrator
fbc52c3f-28ad-4303-a892-8a056630b8f1	BuiltInRole	AppGw for Containers Configuration Manager
4ba50f17-9666-485c-a643-ff00808643f0	BuiltInRole	FHIR SMART User
a001fd3d-188f-4b5d-821b-7da978bf7442	BuiltInRole	Cognitive Services OpenAI Contributor
5e0bd9bd-7b93-4f28-af87-19fc36ad61bd	BuiltInRole	Cognitive Services OpenAI User
36e80216-a7e8-4f42-a7e1-f12c98cbaf8a	BuiltInRole	Impact Reporter
68ff5d27-c7f5-4fa9-a21c-785d0df7bd9e	BuiltInRole	Impact Reader
1afdec4b-e479-420e-99e7-f82237c7c5e6	BuiltInRole	Azure Kubernetes Service Cluster Monitoring User
ad2dd5fb-cd4b-4fd4-a9b6-4fed3630980b	BuiltInRole	ContainerApp Reader
f5819b54-e033-4d82-ac66-4fec3cbf3f4c	BuiltInRole	Azure Connected Machine Resource Manager
189207d4-bb67-4208-a635-b06afe8b2c57	BuiltInRole	SqlDb Migration Role
c4bc862a-3b64-4a35-a021-a380c159b042	BuiltInRole	Bayer Ag Powered Services GDU Solution
ef29765d-0d37-4119-a4f8-f9f9902c9588	BuiltInRole	Bayer Ag Powered Services Imagery Solution
0105a6b0-4bb9-43d2-982a-12806f9faddb	BuiltInRole	Azure Center for SAP solutions Service role for management
6d949e1d-41e2-46e3-8920-c6e4f31a8310	BuiltInRole	Azure Center for SAP solutions Management role
d5a2ae44-610b-4500-93be-660a0c5f5ca6	BuiltInRole	Kubernetes Agentless Operator
f0310ce6-e953-4cf8-b892-fb1c87eaf7f6	BuiltInRole	Azure Usage Billing Data Sender
96062cf7-95ca-4f89-9b9d-2a2aa47356af	BuiltInRole	Azure Container Registry secure supply chain operator service role
1d335eef-eee1-47fe-a9e0-53214eba8872	BuiltInRole	SqlMI Migration Role
a9b99099-ead7-47db-8fcf-072597a61dfa	BuiltInRole	Bayer Ag Powered Services CWUM Solution
ae8036db-e102-405b-a1b9-bae082ea436d	BuiltInRole	SqlVM Migration Role
3f2eb865-5811-4578-b90a-6fc6fa0df8e5	BuiltInRole	Azure Front Door Secret Contributor
0ab34830-df19-4f8c-b84e-aa85b8afa6e8	BuiltInRole	Azure Front Door Domain Contributor
0f99d363-226e-4dca-9920-b807cf8e1a5f	BuiltInRole	Azure Front Door Domain Reader
0db238c4-885e-4c4f-a933-aa2cef684fca	BuiltInRole	Azure Front Door Secret Reader
bfc3b73d-c6ff-45eb-9a5f-40298295bf20	BuiltInRole	LocalRulestacksAdministrator role
d18ad5f3-1baf-4119-b49b-d944edb1f9d0	BuiltInRole	MySQL Backup And Export Operator
bda0d508-adf1-4af0-9c28-88919fc3ae06	BuiltInRole	Azure Stack HCI Administrator
a8835c7d-b5cb-47fa-b6f0-65ea10ce07a2	BuiltInRole	LocalNGFirewallAdministrator role
7392c568-9289-4bde-aaaa-b7131215889d	BuiltInRole	Azure Extension for SQL Server Deployment
d6470a16-71bd-43ab-86b3-6f3a73f4e787	BuiltInRole	Azure Maps Data Read and Batch Role
56328988-075d-4c6a-8766-d93edd6725b6	BuiltInRole	API Management Workspace API Developer
0c34c906-8d99-4cb7-8bb7-33f5b0a1a799	BuiltInRole	API Management Workspace Contributor
ef1c2c96-4a77-49e8-b9a4-6179fe1d2fd2	BuiltInRole	API Management Workspace Reader
73c2c328-d004-4c5e-938c-35c6f5679a1f	BuiltInRole	API Management Workspace API Product Manager
9565a273-41b9-4368-97d2-aeb0c976a9b3	BuiltInRole	API Management Service Workspace API Developer
d59a3e9c-6d52-4a5a-aeed-6bf3cf0e31da	BuiltInRole	API Management Service Workspace API Product Manager
69566ab7-960f-475b-8e7c-b3118f30c6bd	BuiltInRole	Storage File Data Privileged Contributor
b8eda974-7b85-4f76-af95-65846b26df6d	BuiltInRole	Storage File Data Privileged Reader
7eabc9a4-85f7-4f71-b8ab-75daaccc1033	BuiltInRole	Windows 365 Network User
3d55a8f6-4133-418d-8051-facdb1735758	BuiltInRole	Windows365SubscriptionReader
1f135831-5bbe-4924-9016-264044c00788	BuiltInRole	Windows 365 Network Interface Contributor
ffc6bbe0-e443-4c3b-bf54-26581bb2f78e	BuiltInRole	App Compliance Automation Reader
0f37683f-2463-46b6-9ce7-9b788b988ba2	BuiltInRole	App Compliance Automation Administrator
8b9dfcab-4b77-4632-a6df-94bd07820648	BuiltInRole	Azure Sphere Contributor
e9b8712a-cbcf-4ea7-b0f7-e71b803401e6	BuiltInRole	SaaS Hub Contributor
6d994134-994b-4a59-9974-f479f0b227fb	BuiltInRole	Azure Sphere Publisher
c8ae6279-5a0b-4cb2-b3f0-d4d62845742c	BuiltInRole	Azure Sphere Reader
ea01e6af-a1c1-4350-9563-ad00f8c72ec5	BuiltInRole	Azure Machine Learning Workspace Connection Secrets Reader
be1a1ac2-09d3-4261-9e57-a73a6e227f53	BuiltInRole	Procurement Contributor
79b01272-bf9f-4f4c-9517-5506269cf524	BuiltInRole	Cognitive Search Serverless Data Reader (Deprecated)
7ac06ca7-21ca-47e3-a67b-cbd6e6223baf	BuiltInRole	Cognitive Search Serverless Data Contributor (Deprecated)
5e28a61e-8040-49db-b175-bb5b88af6239	BuiltInRole	Community Owner Role
9c1607d1-791d-4c68-885d-c7b7aaff7c8a	BuiltInRole	Firmware Analysis Admin
8b54135c-b56d-4d72-a534-26097cfdc8d8	BuiltInRole	Key Vault Data Access Administrator
1e7ca9b1-60d1-4db8-a914-f2ca1ff27c40	BuiltInRole	Defender for Storage Data Scanner
df2711a6-406d-41cf-b366-b0250bff9ad1	BuiltInRole	Compute Diagnostics Role
fa6cecf6-5db3-4c43-8470-c540bcb4eafa	BuiltInRole	Elastic SAN Network Admin
bba48692-92b0-4667-a9ad-c31c7b334ac2	BuiltInRole	Cognitive Services Usages Reader
c088a766-074b-43ba-90d4-1fb21feae531	BuiltInRole	PostgreSQL Flexible Server Long Term Retention Backup Role
a02f7c31-354d-4106-865a-deedf37fa038	BuiltInRole	Search Parameter Manager
66f75aeb-eabe-4b70-9f1e-c350c4c9ad04	BuiltInRole	Virtual Machine Data Access Administrator (preview)
ad710c24-b039-4e85-a019-deb4a06e8570	BuiltInRole	Logic Apps Standard Contributor (Preview)
b70c96e9-66fe-4c09-b6e7-c98e69c98555	BuiltInRole	Logic Apps Standard Operator (Preview)
4accf36b-2c05-432f-91c8-5c532dff4c73	BuiltInRole	Logic Apps Standard Reader (Preview)
523776ba-4eb2-4600-a3c8-f2dc93da4bdb	BuiltInRole	Logic Apps Standard Developer (Preview)
7b3e853f-ad5d-4fb5-a7b8-56a3581c7037	BuiltInRole	IPAM Pool User
e9c9ed2b-2a99-4071-b2ff-5b113ebf73a1	BuiltInRole	SpatialMapsAccounts Account Owner
0b962ed2-6d56-471c-bd5f-3477d83a7ba4	BuiltInRole	Azure Resource Notifications System Topics Subscriber
90e8b822-3e73-47b5-868a-787dc80c008f	BuiltInRole	Elastic SAN Volume Importer
1c4770c0-34f7-4110-a1ea-a5855cc7a939	BuiltInRole	Elastic SAN Snapshot Exporter
49435da6-99fe-48a5-a235-fc668b9dc04a	BuiltInRole	Community Contributor Role
a12b0b94-b317-4dcd-84a8-502ce99884c6	BuiltInRole	EventGrid TopicSpaces Publisher
4b0f2fd7-60b4-4eca-896f-4435034f8bf5	BuiltInRole	EventGrid TopicSpaces Subscriber
d1a38570-4b05-4d70-b8e4-1100bcf76d12	BuiltInRole	Data Boundary Tenant Administrator
bb6577c4-ea0a-40b2-8962-ea18cb8ecd4e	BuiltInRole	DeID Realtime Data User
8a90fa6b-6997-4a07-8a95-30633a7c97b9	BuiltInRole	DeID Batch Data Owner
b73a14ee-91f5-41b7-bd81-920e12466be9	BuiltInRole	DeID Batch Data Reader
fa0d39e6-28e5-40cf-8521-1eb320653a4c	BuiltInRole	Carbon Optimization Reader
38863829-c2a4-4f8d-b1d2-2e325973ebc7	BuiltInRole	Landing Zone Management Owner
8fe6e843-6d9e-417b-9073-106b048f50bb	BuiltInRole	Landing Zone Management Reader
865ae368-6a45-4bd1-8fbf-0d5151f56fc1	BuiltInRole	Azure Stack HCI Device Management Role
4dae6930-7baf-46f5-909e-0383bc931c46	BuiltInRole	Azure Customer Lockbox Approver for Subscription
7b1f81f9-4196-4058-8aae-762e593270df	BuiltInRole	Azure Resource Bridge Deployment Role
874d1c73-6003-4e60-a13a-cb31ea190a85	BuiltInRole	Azure Stack HCI VM Contributor
64702f94-c441-49e6-a78b-ef80e0188fee	BuiltInRole	Azure AI Developer
4b3fe76c-f777-4d24-a2d7-b027b0f7b273	BuiltInRole	Azure Stack HCI VM Reader
eb960402-bf75-4cc3-8d68-35b34f960f72	BuiltInRole	Deployment Environments Reader
3afb7f49-54cb-416e-8c09-6dc049efa503	BuiltInRole	Azure AI Inference Deployment Operator
1d8c3fe3-8864-474b-8749-01e3783e8157	BuiltInRole	EventGrid Data Contributor
78cbd9e7-9798-4e2e-9b5a-547d9ebb31fb	BuiltInRole	EventGrid Data Receiver
65a14201-8f6c-4c28-bec4-12619c5a9aaa	BuiltInRole	Connected Cluster Managed Identity CheckAccess Reader
c64499e0-74c3-47ad-921c-13865957895c	BuiltInRole	Advisor Reviews Reader
8aac15f0-d885-4138-8afa-bfb5872f7d13	BuiltInRole	Advisor Reviews Contributor
a8d4b70f-0fb9-4f72-b267-b87b2f990aec	BuiltInRole	AgFood Platform Dataset Admin
0f641de8-0b88-4198-bdef-bd8b45ceba96	BuiltInRole	Defender for Storage Scanner Operator
662802e2-50f6-46b0-aed2-e834bacc6d12	BuiltInRole	Azure Front Door Profile Reader
86fede04-b259-4277-8c3e-e26b9865abd8	BuiltInRole	Enclave Reader Role
b5092dac-c796-4349-8681-1a322a31c3f9	BuiltInRole	Azure Kubernetes Service Hybrid Cluster Admin Role
e7037d40-443a-4434-a3fb-8cd202011e1d	BuiltInRole	Azure Kubernetes Service Hybrid Contributor Role
fc3f91a1-40bf-4439-8c46-45edbd83563a	BuiltInRole	Azure Kubernetes Service Hybrid Cluster User Role
3d5f3eff-eb94-473d-91e3-7aac74d6c0bb	BuiltInRole	Enclave Owner Role
e6aadb6b-e64f-41c0-9392-d2bba3bc3ebc	BuiltInRole	Community Reader Role
19feefae-eacc-4106-81fd-ac34c0671f14	BuiltInRole	Enclave Contributor Role
44f0a1a8-6fea-4b35-980a-8ff50c487c97	BuiltInRole	Operator Nexus Key Vault Writer Service Role (Preview)
a316ed6d-1efe-48ac-ac08-f7995a9c26fb	BuiltInRole	Storage Account Encryption Scope Contributor Role
08bbd89e-9f13-488c-ac41-acfcb10c90ab	BuiltInRole	Key Vault Crypto Service Release User
0cd9749a-3aaf-4ae5-8803-bd217705bf3b	BuiltInRole	Kubernetes Runtime Storage Class Contributor Role
609c0c20-e0a0-4a71-b99f-e7e755ac493d	BuiltInRole	Azure Programmable Connectivity Gateway User
db79e9a7-68ee-4b58-9aeb-b90e7c24fcba	BuiltInRole	Key Vault Certificate User
52fd16bd-6ed5-46af-9c40-29cbd7952a29	BuiltInRole	Azure Spring Apps Managed Components Log Reader Role
4301dc2a-25a9-44b0-ae63-3636cf7f2bd2	BuiltInRole	Azure Spring Apps Spring Cloud Gateway Log Reader Role
c7244dfb-f447-457d-b2ba-3999044d1706	BuiltInRole	Azure API Center Data Reader
6593e776-2a30-40f9-8a32-4fe28b77655d	BuiltInRole	Azure Spring Apps Application Configuration Service Log Reader Role
207bcc4b-86a6-4487-9141-d6c1f4c238aa	BuiltInRole	Azure Edge On-Site Deployment Engineer
dfb2f09d-25f8-4558-8986-497084006d7a	BuiltInRole	Azure impact-insight reader
8bb6f106-b146-4ee6-a3f9-b9c5a96e0ae5	BuiltInRole	Defender Kubernetes Agent Operator
8b32b316-c2f5-4ddf-b05b-83dacd2d08b5	BuiltInRole	Azure Red Hat OpenShift Image Registry Operator Role
4436bae4-7702-4c84-919b-c4069ff25ee2	BuiltInRole	Azure Red Hat OpenShift Service Operator Role
be7a6435-15ae-4171-8f30-4a343eff9e8f	BuiltInRole	Azure Red Hat OpenShift Network Operator Role
5b7237c5-45e1-49d6-bc18-a1f62f400748	BuiltInRole	Azure Red Hat OpenShift Storage Operator Role
0336e1d3-7a87-462b-b6db-342b63f7802c	BuiltInRole	Azure Red Hat OpenShift Cluster Ingress Operator Role
0d7aedc0-15fd-4a67-a412-efad370c947e	BuiltInRole	Azure Red Hat OpenShift Azure Files Storage Operator Role
a1f96423-95ce-4224-ab27-4e3dc72facd4	BuiltInRole	Azure Red Hat OpenShift Cloud Controller Manager Role
0358943c-7e01-48ba-8889-02cc51d78637	BuiltInRole	Azure Red Hat OpenShift Machine API Operator Role
5a382001-fe36-41ff-bba4-8bf06bd54da9	BuiltInRole	Azure Sphere Owner
d0f495dc-44ef-4140-aeb0-b89110e6a7c1	BuiltInRole	GroupQuota Reader
e2217c0e-04bb-4724-9580-91cf9871bc01	BuiltInRole	GroupQuota Request Operator
539283cd-c185-4a9a-9503-d35217a1db7b	BuiltInRole	Bayer Ag Powered Services Smart Boundary Solution User Role
8480c0f0-4509-4229-9339-7c10018cb8c4	BuiltInRole	Defender CSPM Storage Scanner Operator
87a87389-f3af-4c43-a694-f6e5efec8582	BuiltInRole	Microsoft Defender for Cloud administrator (preview)
6b534d80-e337-47c4-864f-140f5c7f593d	BuiltInRole	Advisor Recommendations Contributor (Assessments and Reviews)
c9c97b9c-105d-4bb5-a2a7-7d15666c2484	BuiltInRole	GeoCatalog Administrator
b7b8f583-43d0-40ae-b147-6b46f53661c1	BuiltInRole	GeoCatalog Reader
eb5a76d5-50e7-4c33-a449-070e7c9c4cf2	BuiltInRole	Health Bot Reader
f1082fec-a70f-419f-9230-885d2550fb38	BuiltInRole	Health Bot Admin
af854a69-80ce-4ff7-8447-f1118a2e0ca8	BuiltInRole	Health Bot Editor
c20923c5-b089-47a5-bf67-fd89569c4ad9	BuiltInRole	Azure Programmable Connectivity Gateway Dataplane User
b556d68e-0be0-4f35-a333-ad7ee1ce17ea	BuiltInRole	Azure AI Enterprise Network Connection Approver
08d4c71a-cc63-4ce4-a9c8-5dd251b4d619	BuiltInRole	Azure Container Storage Operator
95dd08a6-00bd-4661-84bf-f6726f83a4d0	BuiltInRole	Azure Container Storage Contributor
95de85bd-744d-4664-9dde-11430bc34793	BuiltInRole	Azure Container Storage Owner
5d3f1697-4507-4d08-bb4a-477695db5f82	BuiltInRole	Azure Kubernetes Service Arc Contributor Role
233ca253-b031-42ff-9fba-87ef12d6b55f	BuiltInRole	Azure Kubernetes Service Arc Cluster User Role
b29efa5f-7782-4dc3-9537-4d5bc70a5e9f	BuiltInRole	Azure Kubernetes Service Arc Cluster Admin Role
f54b6d04-23c6-443e-b462-9c16ab7b4a52	BuiltInRole	Backup MUA Operator
c2a970b4-16a7-4a51-8c84-8a8ea6ee0bb8	BuiltInRole	Backup MUA Admin
3d24a3a0-c154-4f6f-a5ed-adc8e01ddb74	BuiltInRole	Savings plan Purchaser
b6ee44de-fe58-4ddc-b5c2-ab174eb23f05	BuiltInRole	CrossConnectionReader
399c3b2b-64c2-4ff1-af34-571db925b068	BuiltInRole	CrossConnectionManager
5e93ba01-8f92-4c7a-b12a-801e3df23824	BuiltInRole	Kubernetes Agent Operator
dd24193f-ef65-44e5-8a7e-6fa6e03f7713	BuiltInRole	Azure API Center Service Contributor
ede9aaa3-4627-494e-be13-4aa7c256148d	BuiltInRole	Azure API Center Compliance Manager
6cba8790-29c5-48e5-bab1-c7541b01cb04	BuiltInRole	Azure API Center Service Reader
e9ce8739-6fa2-4123-a0a2-0ef41a67806f	BuiltInRole	Oracle.Database VmCluster Administrator Built-in Role
4cfdd23b-aece-4fd1-b614-ad3a06c53453	BuiltInRole	Oracle.Database Exadata Infrastructure Administrator Built-in Role
4caf51ec-f9f5-413f-8a94-b9f5fddba66b	BuiltInRole	Oracle Subscriptions Manager Built-in Role
4562aac9-b209-4bd7-a144-6d7f3bb516f4	BuiltInRole	Oracle.Database Owner Built-in Role
b5b192c1-773c-4543-bfb0-6c59254b74a9	BuiltInRole	Bayer Ag Powered Services Historical Weather Data Solution User Role
d623d097-b882-4e1e-a26f-ac60e31065a1	BuiltInRole	Oracle.Database Reader Built-in Role
f27b7598-bc64-41f7-8a44-855ff16326c2	BuiltInRole	Azure Messaging Catalog Data Owner
25211fc6-dc78-40b6-b205-e4ac934fd9fd	BuiltInRole	Azure Spring Apps Application Configuration Service Config File Pattern Reader Role
5d9c6a55-fc0e-4e21-ae6f-f7b095497342	BuiltInRole	Azure Hybrid Database Administrator - Read Only Service Role
30b3bcf2-670a-4bdc-8669-7e0ae0c0dfda	BuiltInRole	Quantum Workspace Owner
c18f9900-27b8-47c7-a8f0-5b3b3d4c2bc2	BuiltInRole	Microsoft Sentinel Business Applications Agent Operator
0fb8eba5-a2bb-4abe-b1c1-49dfad359bb0	BuiltInRole	Azure ContainerApps Session Executor
83ee7727-862c-4213-8ed8-2ce6c5d69a40	BuiltInRole	Microsoft.Edge Winfields federated subscription read access role
ef318e2a-8334-4a05-9e4a-295a196c6a6e	BuiltInRole	Azure Red Hat OpenShift Federated Credential Role
39138f76-04e6-41f0-ba6b-c411b59081a9	BuiltInRole	Bayer Ag Powered Services Crop Id Solution User Role
b67fe603-310e-4889-b9ee-8257d09d353d	BuiltInRole	Scheduled Events Contributor
e82342c9-ac7f-422b-af64-e426d2e12b2d	BuiltInRole	Compute Recommendations Role
91422e52-bb88-4415-bb4a-90f5b71f6dcb	BuiltInRole	Azure Spring Apps Job Execution Instance List Role
b459aa1d-e3c8-436f-ae21-c0531140f43e	BuiltInRole	Azure Spring Apps Job Log Reader Role
a5eb8433-97a5-4a06-80b2-a877e1622c31	BuiltInRole	Nexus Network Fabric Service Writer
05fdd44c-adc6-4aff-981c-61041f0c929a	BuiltInRole	Nexus Network Fabric Service Reader
adb29209-aa1d-457b-a786-c913953d2891	BuiltInRole	Azure Deployment Stack Owner
bf7f8882-3383-422a-806a-6526c631a88a	BuiltInRole	Azure Deployment Stack Contributor
94877a25-7520-40c5-9c42-68e02e4758bd	BuiltInRole	Resource Group Contributor
74252426-c508-480e-9345-4607bbebead4	BuiltInRole	Azure Spring Apps Spring Cloud Config Server Log Reader Role
bfdb9389-c9a5-478a-bb2f-ba9ca092c3c7	BuiltInRole	Container Registry Repository Catalog Lister
b93aa761-3e63-49ed-ac28-beffa264f7ac	BuiltInRole	Container Registry Repository Reader
2a1e307c-b015-4ebd-883e-5b7698a07328	BuiltInRole	Container Registry Repository Writer
2efddaa5-3f1f-4df3-97df-af3f13818f4c	BuiltInRole	Container Registry Repository Contributor
28bf596f-4eb7-45ce-b5bc-6cf482fec137	BuiltInRole	Locks Contributor
78e4b983-1a0b-472e-8b7d-8d770f7c5890	BuiltInRole	DeID Data Owner
8ca4f98b-e38b-4784-9fc0-7aefa970cc5b	BuiltInRole	Kusto Contributor
b1e6a0dd-ea0f-4108-8925-7047693f2cfe	BuiltInRole	CloudTest Leased VM Reader
39fcb0de-8844-4706-b050-c28ddbe3ff83	BuiltInRole	Standby Container Group Pool Contributor
2ccf8795-8983-4912-8036-1c45212c95e8	BuiltInRole	ToolchainOrchestrator Admin Role
c5826735-177b-4a0d-a9a3-d0e4b4bda107	BuiltInRole	ToolchainOrchestrator Viewer Role
85a2d0d9-2eba-4c9c-b355-11c2cc0788ab	BuiltInRole	Compute Gallery Artifacts Publisher
4d8c6f2e-3fd6-4d40-826e-93e3dc4c3fc1	BuiltInRole	ProviderHub Reader
a3ab03bc-5350-42ff-b0d5-00207672db55	BuiltInRole	ProviderHub Contributor
c99c945f-8bd1-4fb1-a903-01460aae6068	BuiltInRole	Azure Stack HCI Connected InfraVMs
dfce8971-25e3-42e3-ba33-6055438e3080	BuiltInRole	VM Restore Operator
0847e196-2fd2-4c2f-a48c-fca6fd030f44	BuiltInRole	HDInsight Cluster Admin
4aa368ec-fba9-4e93-81ed-396b3d461cc5	BuiltInRole	Operator Nexus Compute Contributor Role (Preview)
5d977122-f97e-4b4d-a52f-6b43003ddb4d	BuiltInRole	Azure Container Instances Contributor Role
6cdbb904-5ff3-429d-8169-7d7818b91bd8	BuiltInRole	Connector Reader
b7ef99e8-10af-4a7c-8b15-5b6c352a8378	BuiltInRole	Azure Databases ARM Management Contributor
c20ab07d-648c-4fed-977e-f917d8095dfc	BuiltInRole	Azure Migrate Management Role
8ad4d0ee-9bfb-49e8-93fc-01abb8db6240	BuiltInRole	Transparency Logs Owner
136d308c-0937-4a49-9bd7-edfb42adbffc	BuiltInRole	Disk Encryption Set Operator for Managed Disks
41e04612-9dac-4699-a02b-c82ff2cc3fb5	BuiltInRole	Grafana Limited Viewer
1cbe04d4-3fa3-4559-a406-71584a32cacb	BuiltInRole	Register Contributor
753ceed7-553b-434b-a2e5-78da1a8379d9	BuiltInRole	Elastic Autopilot Contributor
10550c65-218e-4de9-9e52-0200c8ed4748	BuiltInRole	PowerBI Dedicated Contributor
1af232de-e806-426f-8ca1-c36142449755	BuiltInRole	Bayer Ag Powered Services Field Imagery Solution Service Role
cd7709b9-d129-4dc4-ac27-d215a1be1e07	BuiltInRole	Subscription Migrator
a60b64c0-1adf-4051-956a-78f3ae578c7d	BuiltInRole	PostgreSQL Flexible Management Service Contributor (Deprecated)
9295f069-25d0-4f44-bb6a-3da70d11aa00	BuiltInRole	Azure Edge Hardware Center Administrator
b78c5d69-af96-48a3-bf8d-a8b4d589de94	BuiltInRole	Azure AI Administrator
cf7c76d2-98a3-4358-a134-615aa78bf44d	BuiltInRole	Compute Gallery Image Reader
af61e8fc-2633-4b95-bed3-421ad6826515	BuiltInRole	Container Apps SessionPools Reader
57cc5028-e6a7-4284-868d-0611c5923f8d	BuiltInRole	Container Apps ManagedEnvironments Contributor
1b32c00b-7eff-4c22-93e6-93d11d72d2d8	BuiltInRole	Container Apps ManagedEnvironments Reader
358470bc-b998-42bd-ab17-a7e34c199c0f	BuiltInRole	Container Apps Contributor
f7669afb-68b2-44b4-9c5f-6d2a47fddda0	BuiltInRole	Container Apps SessionPools Contributor
edd66693-d32a-450b-997d-0158c03976b0	BuiltInRole	Container Apps Jobs Reader
f3bd1b5c-91fa-40e7-afe7-0c11d331232c	BuiltInRole	Container Apps Operator
4e3d2b60-56ae-4dc6-a233-09c8e5a82e68	BuiltInRole	Container Apps Jobs Contributor
0ad04412-c4d5-4796-b79c-f76d14c8d402	BuiltInRole	Durable Task Data Contributor
1a5682fc-4f12-4b25-927e-e8cfed0c539e	BuiltInRole	KubernetesRuntime Load Balancer Contributor Role
d715fb95-a0f0-4f1c-8be6-5ad2d2767f67	BuiltInRole	AVS Orchestrator Role
db7003cd-07a9-490c-bfa5-23e40314f8d7	BuiltInRole	Service Connector Contributor
cc3c084f-9a2e-4664-b2bc-47a6685a5f99	BuiltInRole	PostgreSQL Flexible Management Contributor
2a740172-0fc2-4039-972c-b31864cd47d6	BuiltInRole	Azure Device Update Agent
2142ea27-02ad-4094-bfea-2dbac6d24934	BuiltInRole	Enclave Approver Role
a68e7c17-0ab2-4c09-9a58-125dae29748c	BuiltInRole	Key Vault Purge Operator
b5b0c71d-aca9-4081-aee2-9b1bb335fc1a	BuiltInRole	Cognitive Services Face Contributor
b9a307c4-5aa3-4b52-ba60-2b17c136cd7b	BuiltInRole	Container Apps Jobs Operator
77be276d-fb44-4f3b-beb5-9bf03c4cd2d3	BuiltInRole	Operator Nexus Owner (Preview)
4e9d0bd4-5aab-4f91-92df-9def33fe287c	BuiltInRole	CloudTest Contributor Role
8d6517c1-e434-405c-9f3f-e0ae65085d76	BuiltInRole	Azure Automanage Contributor
9fc6112f-f48e-4e27-8b09-72a5c94e4ae9	BuiltInRole	Azure Bot Service Contributor Role
175b81b9-6e0d-490a-85e4-0d422273c10c	BuiltInRole	App Configuration Reader
fe86443c-f201-4fc4-9d2a-ac61149fbda0	BuiltInRole	App Configuration Contributor
83f80186-3729-438c-ad2d-39e94d718838	BuiltInRole	Service Fabric Managed Cluster Contributor
577a9874-89fd-4f24-9dbd-b5034d0ad23a	BuiltInRole	Container Registry Data Importer and Data Reader
a35466a1-cfd6-450a-b35e-683fcdf30363	BuiltInRole	Azure Batch Service Orchestration Role
8c87871d-6201-42da-abb1-1c0c985ff71c	BuiltInRole	Microsoft PowerBI Tenant Operations Role
8894c184-eeb6-45cc-868a-1be782055af3	BuiltInRole	MySQL User Data Writer
b5207fd7-42d3-40fa-8cc6-a7e189aef39e	BuiltInRole	MySQL Control Plane
abed7cd4-f358-4384-9d17-def9303b3b53	BuiltInRole	MySQL User Data Reader
b6efc156-f0da-4e90-a50a-8c000140b017	BuiltInRole	Service Fabric Cluster Contributor
6e0c8711-85a0-4490-8365-8ec13c4560b4	BuiltInRole	Stream Analytics Contributor
1dfc38e8-6ce7-447f-807c-029c65262c5f	BuiltInRole	Stream Analytics Reader
80d0d6b0-f522-40a4-8886-a5a11720c375	BuiltInRole	Durable Task Worker
337a31c1-4e14-4ef9-83ed-584bb8d2b70a	BuiltInRole	Fabric Resource Management Administrator
c840cbbc-8508-4228-b700-ec6522a74314	BuiltInRole	Foundational RP Contributor
2718b1f7-eb07-424e-8868-0137541392a1	BuiltInRole	Landing Zone Account Reader
bf2b6809-e9a5-4aea-a6e1-40a9dc8c43a7	BuiltInRole	Landing Zone Account Owner
21bffb94-04c0-4ed0-b676-68bb926e832b	BuiltInRole	Microsoft.Windows365.CloudPcDelegatedMsis Writer User
11076f67-66f6-4be0-8f6b-f0609fd05cc9	BuiltInRole	Azure Batch Account Reader
6aaa78f1-f7de-44ca-8722-c64a23943cae	BuiltInRole	Azure Batch Data Contributor
48e5e92e-a480-4e71-aa9c-2778f4c13781	BuiltInRole	Azure Batch Job Submitter
29fe4964-1e60-436b-bd3a-77fd4c178b3c	BuiltInRole	Azure Batch Account Contributor
0b6ca2e8-2cdc-4bd6-b896-aa3d8c21fc35	BuiltInRole	Defender CSPM Storage Data Scanner
5c2d7e57-b7c2-4d8a-be4f-82afa42c6e95	BuiltInRole	Azure Managed Grafana Workspace Contributor
de754d53-652d-4c75-a67f-1e48d8b49c97	BuiltInRole	Service Group Reader
19c28022-e58e-450d-a464-0b2a53034789	BuiltInRole	Cognitive Services Data Contributor (Preview)
d5adeb5b-107f-4aca-99ea-4e3f4fc008d5	BuiltInRole	Container Apps ConnectedEnvironments Reader
bd80684d-2f5f-4130-892a-0955546282de	BuiltInRole	Azure Kubernetes Fleet Manager RBAC Cluster Reader
1dc4cd5a-de51-4ee4-bc8e-b40e9c17e320	BuiltInRole	Azure Kubernetes Fleet Manager RBAC Cluster Writer
6f4fe6fc-f04f-4d97-8528-8bc18c848dca	BuiltInRole	Container Apps ConnectedEnvironments Contributor
ff478a4e-8633-416e-91bc-ec33ce7c9516	BuiltInRole	Azure Messaging Connectors Owner
3bc748fc-213d-45c1-8d91-9da5725539b9	BuiltInRole	Container Registry Contributor and Data Access Configuration Administrator
7fd69092-c9bc-4b59-9e2e-bca63317e147	BuiltInRole	App Configuration Data SAS User
566f0da3-e2a5-4393-9089-763f8bab8fb6	BuiltInRole	Health Safeguards Data User
```