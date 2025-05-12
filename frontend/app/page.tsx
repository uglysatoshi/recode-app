"use client";

import { useState } from 'react'

export default function Home() {
    const [message, setMessage] = useState<string>('Загрузка...')

    fetch('http://localhost:8080/hello')
        .then((res) => {
            if (!res.ok) throw new Error('Network response was not ok')
            return res.json()
        })
        .then((data) => setMessage(data.message))
        .catch((err) => setMessage('Ошибка: ' + err.message))

    return <h1>{message}</h1>
}