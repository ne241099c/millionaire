import { useState } from 'react';

export default function CountButton({ initialCount = 0 }) {
    const [count, setCount] = useState(initialCount);

    return (
        <button onClick={() => setCount(count + 1)}>
            Count: {count}
        </button>
    );
}