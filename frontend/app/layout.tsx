import React from 'react';

export default function RootLayout({ children }: { children: React.ReactNode }) {
    return (
        <html lang="ru">
        <body className="bg-gray-50">
        <main >
            {children}
        </main>
        </body>
        </html>
    );
}