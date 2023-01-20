package n_able

type Device struct {
	Openapi string `yaml:"openapi"`
	Info    struct {
		Title       string `yaml:"title"`
		Description string `yaml:"description"`
		Version     string `yaml:"version"`
	} `yaml:"info"`
	Servers []struct {
		URL string `yaml:"url"`
	} `yaml:"servers"`
	Paths struct {
		APIDevicesListall struct {
			Get struct {
				Description string `yaml:"description"`
				Responses   struct {
					Num200 struct {
						Description string `yaml:"description"`
						Content     struct {
							ApplicationJSON struct {
								Schema struct {
									Ref string `yaml:"$ref"`
								} `yaml:"schema"`
							} `yaml:"application/json"`
						} `yaml:"content"`
					} `yaml:"200"`
				} `yaml:"responses"`
			} `yaml:"get"`
		} `yaml:"/api/devices/listall"`
		APIDevicesDeviceGUIDMonitorID struct {
			Get struct {
				Description string `yaml:"description"`
				Parameters  []struct {
					In          string `yaml:"in"`
					Name        string `yaml:"name"`
					Required    bool   `yaml:"required"`
					Description string `yaml:"description"`
					Schema      struct {
						Type string `yaml:"type"`
					} `yaml:"schema"`
				} `yaml:"parameters"`
				Responses struct {
					Num200 struct {
						Description string `yaml:"description"`
						Content     struct {
							ApplicationJSON struct {
								Schema struct {
									Ref string `yaml:"$ref"`
								} `yaml:"schema"`
							} `yaml:"application/json"`
						} `yaml:"content"`
					} `yaml:"200"`
				} `yaml:"responses"`
			} `yaml:"get"`
		} `yaml:"/api/devices/{deviceGuid}/{monitorId}"`
	} `yaml:"paths"`
	Components struct {
		Schemas struct {
			Devices struct {
				Type  string `yaml:"type"`
				Items struct {
					Properties struct {
						DeviceID struct {
							Type    string `yaml:"type"`
							Example string `yaml:"example"`
						} `yaml:"deviceId"`
						DeviceName struct {
							Type    string `yaml:"type"`
							Example string `yaml:"example"`
						} `yaml:"deviceName"`
						DeviceGUID struct {
							Type    string `yaml:"type"`
							Example string `yaml:"example"`
						} `yaml:"deviceGuid"`
					} `yaml:"properties"`
				} `yaml:"items"`
			} `yaml:"devices"`
			DeviceMetrics struct {
				Type       string `yaml:"type"`
				Properties struct {
					Num1674145397 struct {
						Type       string `yaml:"type"`
						Properties struct {
							Scantime struct {
								Type    string `yaml:"type"`
								Example struct {
								} `yaml:"example"`
							} `yaml:"scantime"`
							Datadelay struct {
								Type    string `yaml:"type"`
								Example int    `yaml:"example"`
							} `yaml:"datadelay"`
							Errormessage struct {
								Type    string      `yaml:"type"`
								Example interface{} `yaml:"example"`
							} `yaml:"errormessage"`
							State struct {
								Type    string `yaml:"type"`
								Example int    `yaml:"example"`
							} `yaml:"state"`
							Lastupdated struct {
								Type    string `yaml:"type"`
								Example struct {
								} `yaml:"example"`
							} `yaml:"lastupdated"`
							DiskTotal struct {
								Type    string `yaml:"type"`
								Example int64  `yaml:"example"`
							} `yaml:"disk_total"`
							DiskUsed struct {
								Type    string `yaml:"type"`
								Example int    `yaml:"example"`
							} `yaml:"disk_used"`
							DiskFree struct {
								Type    string `yaml:"type"`
								Example int64  `yaml:"example"`
							} `yaml:"disk_free"`
							DiskUsage struct {
								Type    string `yaml:"type"`
								Example int    `yaml:"example"`
							} `yaml:"disk_usage"`
						} `yaml:"properties"`
					} `yaml:"1674145397"`
					Num1674145409 struct {
						Type       string `yaml:"type"`
						Properties struct {
							Scantime struct {
								Type    string `yaml:"type"`
								Example struct {
								} `yaml:"example"`
							} `yaml:"scantime"`
							Datadelay struct {
								Type    string `yaml:"type"`
								Example int    `yaml:"example"`
							} `yaml:"datadelay"`
							Errormessage struct {
								Type    string      `yaml:"type"`
								Example interface{} `yaml:"example"`
							} `yaml:"errormessage"`
							State struct {
								Type    string `yaml:"type"`
								Example int    `yaml:"example"`
							} `yaml:"state"`
							Lastupdated struct {
								Type    string `yaml:"type"`
								Example struct {
								} `yaml:"example"`
							} `yaml:"lastupdated"`
							DiskTotal struct {
								Type    string `yaml:"type"`
								Example int64  `yaml:"example"`
							} `yaml:"disk_total"`
							DiskUsed struct {
								Type    string `yaml:"type"`
								Example int    `yaml:"example"`
							} `yaml:"disk_used"`
							DiskFree struct {
								Type    string `yaml:"type"`
								Example int64  `yaml:"example"`
							} `yaml:"disk_free"`
							DiskUsage struct {
								Type    string `yaml:"type"`
								Example int    `yaml:"example"`
							} `yaml:"disk_usage"`
						} `yaml:"properties"`
					} `yaml:"1674145409"`
				} `yaml:"properties"`
			} `yaml:"deviceMetrics"`
		} `yaml:"schemas"`
	} `yaml:"components"`
}
