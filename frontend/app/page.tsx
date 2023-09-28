import { api } from "@/lib/api";

export default async function Home() {
  const { data: lgtms } = await api.v1.lgtmsList({ limit: 20 });

  return (
    <main>
      <ul>
        {lgtms.map((lgtm) => (
          <li key={lgtm.id}>{lgtm.id}</li>
        ))}
      </ul>
    </main>
  );
}
