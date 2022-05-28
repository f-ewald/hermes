Message statistics
==================
Total messages:    {{ .TotalMessages }}
Received messages: {{ .ReceivedMessages }}
Sent messages:     {{ .SentMessages }}
First message:     {{ .FirstMessage }}
Last message:      {{ .LastMessage }}
Daily Average:     {{ printf "%.2f" .AvgDailyMessages }}
Monthly Average:   <Not available>
Yearly Average:    <Not available>
Chats:             {{ .Chats }}