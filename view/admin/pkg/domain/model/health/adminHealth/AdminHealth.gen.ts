import type { AdminHealthEnum } from "~/pkg/domain/model/health/adminHealth/AdminHealthEnum.gen";

export type AdminHealth = {
  healthId: number;
  name: string;
  adminHealthEnum: AdminHealthEnum;
};
