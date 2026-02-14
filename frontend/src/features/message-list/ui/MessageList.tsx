'use client';

import { useEffect, useState } from 'react';
import { Message, messageApi } from '@/entities/message';

export const MessageList = () => {
    const [messages, setMessages] = useState<Message[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);
    const [newMessage, setNewMessage] = useState('');
    const [editingId, setEditingId] = useState<number | null>(null);
    const [editContent, setEditContent] = useState('');

    const fetchMessages = () => {
        setLoading(true);
        messageApi.getMessages()
            .then(setMessages)
            .catch((err) => setError('Failed to load messages'))
            .finally(() => setLoading(false));
    };

    useEffect(() => {
        fetchMessages();
    }, []);

    const handleCreate = async (e: React.FormEvent) => {
        e.preventDefault();
        if (!newMessage.trim()) return;
        try {
            await messageApi.createMessage(newMessage, 'User');
            setNewMessage('');
            fetchMessages();
        } catch (err) {
            setError('Failed to create message');
        }
    };

    const handleUpdate = async (id: number) => {
        try {
            await messageApi.updateMessage(id, editContent);
            setEditingId(null);
            fetchMessages();
        } catch (err) {
            setError('Failed to update message');
        }
    };

    const handleDelete = async (id: number) => {
        if (!confirm('Are you sure?')) return;
        try {
            await messageApi.deleteMessage(id);
            fetchMessages();
        } catch (err) {
            setError('Failed to delete message');
        }
    };

    if (loading && messages.length === 0) return <div>Loading messages...</div>;
    if (error) return <div className="text-red-500 font-bold">{error}</div>;

    return (
        <div className="mt-4">
            <form onSubmit={handleCreate} className="mb-6 flex gap-2">
                <input
                    type="text"
                    value={newMessage}
                    onChange={(e) => setNewMessage(e.target.value)}
                    placeholder="New Message..."
                    className="flex-1 p-2 border-2 border-black font-mono focus:outline-none"
                />
                <button type="submit" className="mono-button">
                    POST
                </button>
            </form>

            {messages.length === 0 ? (
                <p>No messages found.</p>
            ) : (
                <ul className="space-y-4">
                    {messages.map((msg) => (
                        <li key={msg.id} className="p-4 border-2 border-black relative group">
                            <div className="absolute top-2 right-2 flex gap-2 opacity-0 group-hover:opacity-100 transition-opacity">
                                <button
                                    onClick={() => {
                                        setEditingId(msg.id);
                                        setEditContent(msg.content);
                                    }}
                                    className="text-xs uppercase hover:underline"
                                >
                                    Edit
                                </button>
                                <button
                                    onClick={() => handleDelete(msg.id)}
                                    className="text-xs uppercase text-red-600 hover:underline"
                                >
                                    Delete
                                </button>
                            </div>

                            <div className="font-bold mb-1 uppercase text-sm border-b border-black inline-block">
                                {msg.sender}
                            </div>

                            {editingId === msg.id ? (
                                <div className="mt-2 flex gap-2">
                                    <input
                                        type="text"
                                        value={editContent}
                                        onChange={(e) => setEditContent(e.target.value)}
                                        className="flex-1 p-1 border border-black font-mono text-sm"
                                    />
                                    <button onClick={() => handleUpdate(msg.id)} className="text-xs border border-black px-2 hover:bg-black hover:text-white">Save</button>
                                    <button onClick={() => setEditingId(null)} className="text-xs px-2 hover:underline">Cancel</button>
                                </div>
                            ) : (
                                <div className="mt-2">{msg.content}</div>
                            )}
                        </li>
                    ))}
                </ul>
            )}
        </div>
    );
};
