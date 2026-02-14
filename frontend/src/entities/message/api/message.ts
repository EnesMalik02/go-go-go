import { api } from '@/shared/api/base';
import { Message } from '../model/types';

export const messageApi = {
  getMessages: async (): Promise<Message[]> => {
    const response = await api.get<Message[]>('/api/v1/messages');
    return response.data;
  },
  createMessage: async (content: string, sender: string): Promise<Message> => {
    const response = await api.post<Message>('/api/v1/messages', { content, sender });
    return response.data;
  },
  updateMessage: async (id: number, content: string): Promise<Message> => {
    const response = await api.put<Message>(`/api/v1/messages/${id}`, { content });
    return response.data;
  },
  deleteMessage: async (id: number): Promise<void> => {
    await api.delete(`/api/v1/messages/${id}`);
  },
};
