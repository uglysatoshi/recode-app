"use client"

import "@/app/globals.css"
import React, { ReactNode } from "react"
import {ThemeProvider} from "@/components/theme-provider";


export default function RootLayout({ children }: { children: ReactNode }) {

    return (
        <html lang="ru">
            <body className="flex min-h-screen">
                <main className="flex-1 p-4 overflow-auto">
                    <ThemeProvider
                        attribute="class"
                        defaultTheme="system"
                        enableSystem
                        disableTransitionOnChange
                    >
                        {children}
                    </ThemeProvider>
                </main>
            </body>
        </html>
    )
}