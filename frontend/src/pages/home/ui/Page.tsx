import { Header } from '@/widgets/header';
import { UserList } from '@/features/user-list';
import { MessageList } from '@/features/message-list';

export const HomePage = () => {
    return (
        <div className="min-h-screen bg-white text-black font-mono">
            <Header />
            <main className="container mx-auto p-8">
                <h1 className="text-4xl font-bold mb-8 uppercase tracking-widest border-b-4 border-black pb-2">
                    Dashboard
                </h1>
                <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
                    <section>
                        <h2 className="text-2xl font-bold mb-4 uppercase border-b-2 border-black inline-block">Users</h2>
                        <UserList />
                    </section>
                    <section>
                        <h2 className="text-2xl font-bold mb-4 uppercase border-b-2 border-black inline-block">Messages</h2>
                        <MessageList />
                    </section>
                </div>
            </main>
        </div>
    );
};
