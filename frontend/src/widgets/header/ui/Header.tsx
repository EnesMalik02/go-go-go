import Link from 'next/link';

export const Header = () => {
    return (
        <header className="flex justify-between items-center p-4 bg-white dark:bg-black border-b">
            <div className="font-bold text-lg">My FSD App</div>
            <nav className="flex gap-4">
                <Link href="/" className="hover:underline">Home</Link>
                <Link href="/about" className="hover:underline">About</Link>
            </nav>
        </header>
    );
};
