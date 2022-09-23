<!DOCTYPE html>
<html>
<head>
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/css2?family=Roboto:wght@400;700&display=swap" rel="stylesheet">
    <style>
        body {
            font-family: 'Roboto', sans-serif;
        }

        a {
            color: #1565C0;
        }

        table {
            max-width: 1024px;
            margin-right: auto;
            margin-left: auto;
            border-collapse: collapse;
            width: 100%;
        }

        td,
        th {
            text-align: left;
            padding-top: 16px;
            padding-bottom: 16px;
            border-bottom: 1px solid #E0E0E0;
        }

        th {
            font-weight: 700;
        }
    </style>
</head>
<body>
<table>
    <thead>
    <tr>
        <th style="width: 25%">Package</th>
        <th style="width: 50%">Source</th>
        <th style="width: 25%">Documentation</th>
    </tr>
    </thead>
    <tbody>
    {{ range $pkg := . }}
        <tr>
            <td>{{ $pkg.Name }}</td>
            <td>
                <a href="{{ $pkg.Git }}">{{ clean $pkg.Git }}</a>
            </td>
            <td>
                {{ $url := module $pkg }}
                <a href="https://pkg.go.dev/{{ $url }}"><img src="https://pkg.go.dev/badge/{{ $url }}.svg" alt="Go Reference"></a>
            </td>
        </tr>
    {{ end }}
    </tbody>
</table>
</body>
</html>