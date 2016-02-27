package main

import (
	seed "../../"
)

func main() {
	en := "eyJzaWduYXR1cmVzIjp7ImF0dHJpYnV0ZXMiOiI1NjU2Mjc5YmVjYmRjZjJhMjNkMDllZDhlOGQyYjY1NWI5Y2EyYjllNzcyYTlhY2QwY2Y2YzExMTQ1YjVmMjM2YTdkNDQ0MGNkMjY1ODIwMmQwYjVhMWYyZmM5NDJjYmNlMDAyMDdkODg0NzlhNzE1ZTdiM2Y1MDE2YjIzYTZlNjJhMjhiNDBiMTRkYmQxYzBjOTMxNzI1Y2ExZmJiZjJmNzgzMzgxMzQ2YzYxNjEzM2E2ZDFiOWU5MzcxYmUxNWRlY2QzMTJiZDQxMGY1MDU2ZGJlYzU1NGE5Yzg1MmFjMzAzY2VlMWRhNjdmNzFlMjk3NjA3OGNkODdjMjcyYmZlIiwiZW1iZWRzIjoiMzI0ZTczZTkyNGFjYjkxMzliMTkxOWIxY2RjMTkyZjE3NTEwMTdjZmRjMzg0YzY1ZTcwNjBkYjA5OWY0NDNjNTFiOTE3NTVhM2RhYmMyZWUwZjkxYjI3MDZkZjc3ZDc3NmUzODg0MmEzOGFmODRmODFjMDVlOGIwYzQ1NmMzMGFkNDE5ODc4ZjFlYzRmM2I0YmZhZjJiYjVkNjkwOTM1ZWY1MTdhODNiMTAzN2IyNmZjMGE3ZDdhMTUyMTllMzRkMzYyNjc3YjA3MTM0OTAxYjFlMWIzYTNkZGIzNTcyZmJhZjdiMmZkYzlmZDNmZDJlMzYxZTNkNmRiZDdjODRhNCIsIm1ldGEiOiIzMjRlNzNlOTI0YWNiOTEzOWIxOTE5YjFjZGMxOTJmMTc1MTAxN2NmZGMzODRjNjVlNzA2MGRiMDk5ZjQ0M2M1MWI5MTc1NWEzZGFiYzJlZTBmOTFiMjcwNmRmNzdkNzc2ZTM4ODQyYTM4YWY4NGY4MWMwNWU4YjBjNDU2YzMwYWQ0MTk4NzhmMWVjNGYzYjRiZmFmMmJiNWQ2OTA5MzVlZjUxN2E4M2IxMDM3YjI2ZmMwYTdkN2ExNTIxOWUzNGQzNjI2NzdiMDcxMzQ5MDFiMWUxYjNhM2RkYjM1NzJmYmFmN2IyZmRjOWZkM2ZkMmUzNjFlM2Q2ZGJkN2M4NGE0Iiwib3duZXJzaGlwIjoiOTUzNDc3ODQwNTEyODdmOWZlMDQyZjkwYjQ1MmYyZjY4YzlmNzhiYjkyMGNjNzI0NzAwNWY2NDNjMTcyODUyMzlmZGIwMjRiYmUwMjIwNzY5N2EzZTRjYzQ0OTQ5NTQyM2FlN2ZlNGYxOTlkZWIzN2Q3NjdkMjI2Y2FmZTlmZTUwYzk5MWZiNzk3NzY3OWVmZTliNDEzYTk1YzdkZmQxMjAzOTQ5ZWQyNmYyZDEyOWVkYzdlODYxY2U5MDY5ZGU3NmI4MzU4OTNmMTYzODBmZDFjZDk2Nzk2YTU1ZWNhNjA0ZTc4NzY1OGQ0MmE0MGNlZTc3YmNjZTE2NjFkMzVjMyJ9LCJtZXRhIjp7ImNyZWF0ZWRfYXQiOjEuNDU2MDUwOTczZSswOSwic2VlZF9pZCI6ImUwYjVhODJhNmVmMzYxOTM3MTljZWI3NmNhNTQ4YzNjYTkzMGFiZDIiLCJzZWVkX3R5cGUiOiJjdXJyZW5jeSJ9LCJvd25lcnNoaXAiOnsic29sZSI6eyJhZGRyZXNzX2lkIjoiNTZjMWYzZDFmNjhkYWE0NTg0MDAwMDAxIn0sInR5cGUiOiJzb2xlIn0sImVtYmVkcyI6W3siYXR0cmlidXRlcyI6eyJzdHVmZiI6InN0dWZmIn0sImVtYmVkcyI6bnVsbCwibWV0YSI6eyJjcmVhdGVkX2F0IjoxLjQ1NjA1MDk3M2UrMDksInNlZWRfaWQiOiJlMGI1YTgyYTZlZjM2MTkzNzE5Y2ViNzZjYTU0OGMzY2E5MzBhYmQyIiwic2VlZF90eXBlIjoiY3VycmVuY3kifSwib3duZXJzaGlwIjp7InNvbGUiOnsiYWRkcmVzc19pZCI6IjU2YzFmM2QxZjY4ZGFhNDU4NDAwMDAwMSJ9LCJ0eXBlIjoic29sZSJ9LCJzaWduYXR1cmVzIjp7ImF0dHJpYnV0ZXMiOiJhYmNkZSIsIm1ldGEiOiIzMjRlNzNlOTI0YWNiOTEzOWIxOTE5YjFjZGMxOTJmMTc1MTAxN2NmZGMzODRjNjVlNzA2MGRiMDk5ZjQ0M2M1MWI5MTc1NWEzZGFiYzJlZTBmOTFiMjcwNmRmNzdkNzc2ZTM4ODQyYTM4YWY4NGY4MWMwNWU4YjBjNDU2YzMwYWQ0MTk4NzhmMWVjNGYzYjRiZmFmMmJiNWQ2OTA5MzVlZjUxN2E4M2IxMDM3YjI2ZmMwYTdkN2ExNTIxOWUzNGQzNjI2NzdiMDcxMzQ5MDFiMWUxYjNhM2RkYjM1NzJmYmFmN2IyZmRjOWZkM2ZkMmUzNjFlM2Q2ZGJkN2M4NGE0Iiwib3duZXJzaGlwIjoiYWEifX1dLCJhdHRyaWJ1dGVzIjp7InN0dWZmIjoic3R1ZmYifX0="
	sh, err := seed.Load(en)
	if err != nil {
		seed.Println("Error loading")
		return
	}
	seed.Println(sh.JSON())
}