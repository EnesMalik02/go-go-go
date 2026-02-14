import { Header } from '@/widgets/header';
import { UserList } from '@/features/user-list';

export const HomePage = () => {
    return (
        <div className="min-h-screen bg-gray-50 dark:bg-gray-900 text-gray-900 dark:text-gray-100">
            <Header />
            <main className="container mx-auto p-4">
                <h1 className="text-2xl font-bold mb-6">Welcome to My FSD App</h1>
                <UserList />
            </main>
        </div>
    );
};
