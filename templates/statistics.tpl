Message statistics
==================
Total messages:    {{ .TotalMessages }}
Received messages: {{ .ReceivedMessages }}
Sent messages:     {{ .SentMessages }}
First message:     {{ .FirstMessage }}
Last message:      {{ .LastMessage }}
Daily Average:     {{ printf "%.2f" .AvgDailyMessages }}
Monthly Average:   {{ printf "%.2f" .AvgMonthlyMessages }}
Yearly Average:    {{ printf "%.2f" .AvgYearlyMessages }}
Chats:             {{ .Chats }}