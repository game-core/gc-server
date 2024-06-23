import type {AdminHealth} from "~/pkg/domain/health/adminHealth/AdminHealth"
import type {CommonHealth} from "~/pkg/domain/health/commonHealth/CommonHealth"
import type {MasterHealth} from "~/pkg/domain/health/masterHealth/MasterHealth"

export type HealthCheckResponse = {
    adminHealth: AdminHealth;
    commonHealth: CommonHealth;
    masterHealth: MasterHealth;
}
