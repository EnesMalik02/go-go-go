import { api } from '@/shared/api/base';
import { User } from '../model/types';

export const userApi = {
  getUsers: async (): Promise<User[]> => {
    const response = await api.get<User[]>('/api/v1/users');
    return response.data;
  },
};
