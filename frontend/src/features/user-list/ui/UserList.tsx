'use client';

import { useEffect, useState } from 'react';
import { User, userApi } from '@/entities/user';

export const UserList = () => {
    const [users, setUsers] = useState<User[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        userApi.getUsers()
            .then(setUsers)
            .catch((err) => setError('Failed to load users'))
            .finally(() => setLoading(false));
    }, []);

    if (loading) return <div>Loading users...</div>;
    if (error) return <div className="text-red-500">{error}</div>;

    return (
        <div className="p-4 border rounded-lg shadow-sm">
            <h2 className="text-xl font-bold mb-4">Users</h2>
            {users.length === 0 ? (
                <p>No users found.</p>
            ) : (
                <ul className="space-y-2">
                    {users.map((user) => (
                        <li key={user.id} className="p-2 bg-gray-50 rounded">
                            <div className="font-semibold">{user.name}</div>
                            <div className="text-sm text-gray-500">{user.email}</div>
                        </li>
                    ))}
                </ul>
            )}
        </div>
    );
};
