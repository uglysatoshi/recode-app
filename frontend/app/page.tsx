"use client";

import { useEffect, useState } from 'react'

export default function Home() {
    const [message, setMessage] = useState<string>('Загрузка...')

    useEffect(() => {
        fetch('http://localhost:8080')
            .then((res) => res.json())
            .then((data) => setMessage(data.message))
            .catch((err) => {
                console.error(err)
                setMessage('Ошибка: ' + err.message)
            })
    }, [])

    return <h1>{message}</h1>
}