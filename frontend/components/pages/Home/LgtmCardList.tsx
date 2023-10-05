import React, { useCallback, useState } from "react";
import LgtmCard, { LgtmCardProps } from "./LgtmCard";
import ReportForm from "./ReportForm";

export type LgtmCardListProps = {
  lgtmIds: string[];
} & Omit<LgtmCardProps, "lgtmId" | "onStartReport">;

export default function LgtmCardList({
  lgtmIds,
  ...cardProps
}: LgtmCardListProps) {
  const [reportingLgtmId, setReportingLgtmId] = useState<string | null>(null);

  const handleStartReport = useCallback((id: string) => {
    setReportingLgtmId(id);
  }, []);

  const handleCloseReportForm = useCallback(() => {
    setReportingLgtmId(null);
  }, []);

  return (
    <>
      <ReportForm lgtmId={reportingLgtmId} onClose={handleCloseReportForm} />

      <ul className="grid grid-cols-2 gap-4 sm:grid-cols-3 md:grid-cols-4">
        {lgtmIds.map((id) => (
          <li key={id}>
            <LgtmCard
              className="h-full"
              lgtmId={id}
              onStartReport={handleStartReport}
              {...cardProps}
            />
          </li>
        ))}
      </ul>
    </>
  );
}
