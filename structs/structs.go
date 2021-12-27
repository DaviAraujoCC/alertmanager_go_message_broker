package structs

type Webhook struct {
	Receiver string `json: "receiver"`
	Status   string `json: "status"`

	Alerts struct {
		Status string `json: "status"`
		Labels struct {
			Alertname string `json: "alertname"`
			Instance  string `json: "instance"`
			Job       string `json: "job"`
			Severity  string `json: "severity"`
			Team      string `json: "team"`
		}
		Annotations struct {
			Dashboard   string `json: "dashboard"`
			Description string `json: "description"`
		}
		StartsAt     string `json: "startsAt"`
		EndsAt       string `json: "endsAt"`
		GeneratorURL string `json: "generatorURL"`
	}

	GroupLabels struct {
		Alertname string `json: "alertname"`
		Job       string `json: "job"`
	}

	CommonLabels struct {
		Alertname string `json: "alertname"`
		Dc        string `json: "dc"`
		Instance  string `json: "instance"`
		Job       string `json: "job"`
	}

	CommonAnnotations struct {
		Description string `json: "description"`
	}
	ExternalURL string `json: "externalURL"`
	Version     string `json: "version"`
	GroupKey    string `json: "groupKey"`
}

type Endpoint struct {
	Id        string
	Url       string
	Endpoint  string
	Alertname string
}

type RequestCreateEndpoint struct {
	Action    string `json: "action"`
	Url       string `json: "url"`
	Endpoint  string `json: "endpoint"`
	ID        string `json: "id"`
	Alertname string `json: "alertname"`
}

// Create payloads here

type Payload_test struct {
	Alertname   string `json: "alertname"`
	Instance    string `json: "instance"`
	Job         string `json: "job"`
	Severity    string `json: "severity"`
	Status      string `json: "status"`
	Description string `json: "description"`
}
