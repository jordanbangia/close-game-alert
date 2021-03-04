package espn

type EspnScoreboard struct {
	// Leagues []Leagues `json:"leagues"`
	// Season  Season    `json:"season"`
	// Day     Day       `json:"day"`
	Events []Event `json:"events"`
}

type Event struct {
	ID        string `json:"id"`
	UID       string `json:"uid"`
	Date      string `json:"date"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
	Season    struct {
		Year int    `json:"year"`
		Type int    `json:"type"`
		Slug string `json:"slug"`
	} `json:"season"`
	Competitions []struct {
		ID         string `json:"id"`
		UID        string `json:"uid"`
		Date       string `json:"date"`
		Attendance int    `json:"attendance"`
		Type       struct {
			ID           string `json:"id"`
			Abbreviation string `json:"abbreviation"`
		} `json:"type"`
		TimeValid             bool `json:"timeValid"`
		NeutralSite           bool `json:"neutralSite"`
		ConferenceCompetition bool `json:"conferenceCompetition"`
		Recent                bool `json:"recent"`
		Venue                 struct {
			ID       string `json:"id"`
			FullName string `json:"fullName"`
			Address  struct {
				City  string `json:"city"`
				State string `json:"state"`
			} `json:"address"`
			Capacity int  `json:"capacity"`
			Indoor   bool `json:"indoor"`
		} `json:"venue"`
		Competitors []struct {
			ID       string `json:"id"`
			UID      string `json:"uid"`
			Type     string `json:"type"`
			Order    int    `json:"order"`
			HomeAway string `json:"homeAway"`
			Winner   bool   `json:"winner"`
			Team     struct {
				ID               string `json:"id"`
				UID              string `json:"uid"`
				Location         string `json:"location"`
				Name             string `json:"name"`
				Abbreviation     string `json:"abbreviation"`
				DisplayName      string `json:"displayName"`
				ShortDisplayName string `json:"shortDisplayName"`
				Color            string `json:"color"`
				AlternateColor   string `json:"alternateColor"`
				IsActive         bool   `json:"isActive"`
				Venue            struct {
					ID string `json:"id"`
				} `json:"venue"`
				Links []struct {
					Rel        []string `json:"rel"`
					Href       string   `json:"href"`
					Text       string   `json:"text"`
					IsExternal bool     `json:"isExternal"`
					IsPremium  bool     `json:"isPremium"`
				} `json:"links"`
				Logo string `json:"logo"`
			} `json:"team"`
			Score      string `json:"score"`
			Linescores []struct {
				Value float64 `json:"value"`
			} `json:"linescores"`
			Statistics []struct {
				Name         string `json:"name"`
				Abbreviation string `json:"abbreviation"`
				DisplayValue string `json:"displayValue"`
			} `json:"statistics"`
			Leaders []struct {
				Name             string `json:"name"`
				DisplayName      string `json:"displayName"`
				ShortDisplayName string `json:"shortDisplayName"`
				Abbreviation     string `json:"abbreviation"`
				Leaders          []struct {
					DisplayValue string  `json:"displayValue"`
					Value        float64 `json:"value"`
					Athlete      struct {
						ID          string `json:"id"`
						FullName    string `json:"fullName"`
						DisplayName string `json:"displayName"`
						ShortName   string `json:"shortName"`
						Links       []struct {
							Rel  []string `json:"rel"`
							Href string   `json:"href"`
						} `json:"links"`
						Headshot string `json:"headshot"`
						Jersey   string `json:"jersey"`
						Position struct {
							Abbreviation string `json:"abbreviation"`
						} `json:"position"`
						Team struct {
							ID string `json:"id"`
						} `json:"team"`
						Active bool `json:"active"`
					} `json:"athlete"`
					Team struct {
						ID string `json:"id"`
					} `json:"team"`
				} `json:"leaders"`
			} `json:"leaders"`
			Records []struct {
				Name         string `json:"name"`
				Abbreviation string `json:"abbreviation,omitempty"`
				Type         string `json:"type"`
				Summary      string `json:"summary"`
			} `json:"records"`
		} `json:"competitors"`
		Notes  []interface{} `json:"notes"`
		Status struct {
			Clock        float64 `json:"clock"`
			DisplayClock string  `json:"displayClock"`
			Period       int     `json:"period"`
			Type         struct {
				ID          string `json:"id"`
				Name        string `json:"name"`
				State       string `json:"state"`
				Completed   bool   `json:"completed"`
				Description string `json:"description"`
				Detail      string `json:"detail"`
				ShortDetail string `json:"shortDetail"`
			} `json:"type"`
		} `json:"status"`
		Broadcasts []struct {
			Market string   `json:"market"`
			Names  []string `json:"names"`
		} `json:"broadcasts"`
		StartDate     string `json:"startDate"`
		GeoBroadcasts []struct {
			Type struct {
				ID        string `json:"id"`
				ShortName string `json:"shortName"`
			} `json:"type"`
			Market struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"market"`
			Media struct {
				ShortName string `json:"shortName"`
			} `json:"media"`
			Lang   string `json:"lang"`
			Region string `json:"region"`
		} `json:"geoBroadcasts"`
		Headlines []struct {
			Description   string `json:"description"`
			Type          string `json:"type"`
			ShortLinkText string `json:"shortLinkText"`
			Video         []struct {
				ID        int    `json:"id"`
				Source    string `json:"source"`
				Headline  string `json:"headline"`
				Thumbnail string `json:"thumbnail"`
				Duration  int    `json:"duration"`
				Tracking  struct {
					SportName    string `json:"sportName"`
					LeagueName   string `json:"leagueName"`
					CoverageType string `json:"coverageType"`
					TrackingName string `json:"trackingName"`
					TrackingID   string `json:"trackingId"`
				} `json:"tracking"`
				DeviceRestrictions struct {
					Type    string   `json:"type"`
					Devices []string `json:"devices"`
				} `json:"deviceRestrictions"`
				GeoRestrictions struct {
					Type      string   `json:"type"`
					Countries []string `json:"countries"`
				} `json:"geoRestrictions"`
				Links struct {
					API struct {
						Self struct {
							Href string `json:"href"`
						} `json:"self"`
						Artwork struct {
							Href string `json:"href"`
						} `json:"artwork"`
					} `json:"api"`
					Web struct {
						Href  string `json:"href"`
						Short struct {
							Href string `json:"href"`
						} `json:"short"`
						Self struct {
							Href string `json:"href"`
						} `json:"self"`
					} `json:"web"`
					Source struct {
						Mezzanine struct {
							Href string `json:"href"`
						} `json:"mezzanine"`
						Flash struct {
							Href string `json:"href"`
						} `json:"flash"`
						Hds struct {
							Href string `json:"href"`
						} `json:"hds"`
						HLS struct {
							Href string `json:"href"`
							HD   struct {
								Href string `json:"href"`
							} `json:"HD"`
						} `json:"HLS"`
						HD struct {
							Href string `json:"href"`
						} `json:"HD"`
						Full struct {
							Href string `json:"href"`
						} `json:"full"`
						Href string `json:"href"`
					} `json:"source"`
					Mobile struct {
						Alert struct {
							Href string `json:"href"`
						} `json:"alert"`
						Source struct {
							Href string `json:"href"`
						} `json:"source"`
						Href      string `json:"href"`
						Streaming struct {
							Href string `json:"href"`
						} `json:"streaming"`
						ProgressiveDownload struct {
							Href string `json:"href"`
						} `json:"progressiveDownload"`
					} `json:"mobile"`
				} `json:"links"`
			} `json:"video"`
		} `json:"headlines"`
	} `json:"competitions"`
	Links []struct {
		Language   string   `json:"language"`
		Rel        []string `json:"rel"`
		Href       string   `json:"href"`
		Text       string   `json:"text"`
		ShortText  string   `json:"shortText"`
		IsExternal bool     `json:"isExternal"`
		IsPremium  bool     `json:"isPremium"`
	} `json:"links"`
	Status struct {
		Clock        float64 `json:"clock"`
		DisplayClock string  `json:"displayClock"`
		Period       int     `json:"period"`
		Type         struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			State       string `json:"state"`
			Completed   bool   `json:"completed"`
			Description string `json:"description"`
			Detail      string `json:"detail"`
			ShortDetail string `json:"shortDetail"`
		} `json:"type"`
	} `json:"status"`
}
