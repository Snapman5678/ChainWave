'use client';

export default function Error({
  // uncomment later error,
  reset,
}: {
  error: Error;
  reset: () => void;
}) {
  return (
    <div className="min-h-screen bg-gray-50 flex items-center justify-center">
      <div className="text-center">
        <h2 className="text-lg font-semibold">Something went wrong!</h2>
        <button
          onClick={() => reset()}
          className="mt-4 px-4 py-2 bg-blue-500 text-white rounded-md"
        >
          Try again
        </button>
      </div>
    </div>
  );
}