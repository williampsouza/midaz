package constant

import (
	"errors"
)

/* List of errors that can be returned.
 * For more details, refer to the API documentation: https://docs.midaz.io/midaz/api-reference/resources/errors-list
 */
var (
	ErrDuplicateLedger                          = errors.New("0001")
	ErrLedgerNameConflict                       = errors.New("0002")
	ErrAssetNameOrCodeDuplicate                 = errors.New("0003")
	ErrCodeUppercaseRequirement                 = errors.New("0004")
	ErrCurrencyCodeStandardCompliance           = errors.New("0005")
	ErrUnmodifiableField                        = errors.New("0006")
	ErrEntityNotFound                           = errors.New("0007")
	ErrActionNotPermitted                       = errors.New("0008")
	ErrMissingFieldsInRequest                   = errors.New("0009")
	ErrAccountTypeImmutable                     = errors.New("0010")
	ErrInactiveAccountType                      = errors.New("0011")
	ErrAccountBalanceDeletion                   = errors.New("0012")
	ErrResourceAlreadyDeleted                   = errors.New("0013")
	ErrSegmentIDInactive                        = errors.New("0014")
	ErrDuplicateSegmentName                     = errors.New("0015")
	ErrBalanceRemainingDeletion                 = errors.New("0016")
	ErrInvalidScriptFormat                      = errors.New("0017")
	ErrInsufficientFunds                        = errors.New("0018")
	ErrAccountIneligibility                     = errors.New("0019")
	ErrAliasUnavailability                      = errors.New("0020")
	ErrParentTransactionIDNotFound              = errors.New("0021")
	ErrImmutableField                           = errors.New("0022")
	ErrTransactionTimingRestriction             = errors.New("0023")
	ErrAccountStatusTransactionRestriction      = errors.New("0024")
	ErrInsufficientAccountBalance               = errors.New("0025")
	ErrTransactionMethodRestriction             = errors.New("0026")
	ErrDuplicateTransactionTemplateCode         = errors.New("0027")
	ErrDuplicateAssetPair                       = errors.New("0028")
	ErrInvalidParentAccountID                   = errors.New("0029")
	ErrMismatchedAssetCode                      = errors.New("0030")
	ErrChartTypeNotFound                        = errors.New("0031")
	ErrInvalidCountryCode                       = errors.New("0032")
	ErrInvalidCodeFormat                        = errors.New("0033")
	ErrAssetCodeNotFound                        = errors.New("0034")
	ErrPortfolioIDNotFound                      = errors.New("0035")
	ErrSegmentIDNotFound                        = errors.New("0036")
	ErrLedgerIDNotFound                         = errors.New("0037")
	ErrOrganizationIDNotFound                   = errors.New("0038")
	ErrParentOrganizationIDNotFound             = errors.New("0039")
	ErrInvalidType                              = errors.New("0040")
	ErrTokenMissing                             = errors.New("0041")
	ErrInvalidToken                             = errors.New("0042")
	ErrInsufficientPrivileges                   = errors.New("0043")
	ErrPermissionEnforcement                    = errors.New("0044")
	ErrJWKFetch                                 = errors.New("0045")
	ErrInternalServer                           = errors.New("0046")
	ErrBadRequest                               = errors.New("0047")
	ErrInvalidDSLFileFormat                     = errors.New("0048")
	ErrEmptyDSLFile                             = errors.New("0049")
	ErrMetadataKeyLengthExceeded                = errors.New("0050")
	ErrMetadataValueLengthExceeded              = errors.New("0051")
	ErrAccountIDNotFound                        = errors.New("0052")
	ErrUnexpectedFieldsInTheRequest             = errors.New("0053")
	ErrIDsNotFoundForAccounts                   = errors.New("0054")
	ErrAssetIDNotFound                          = errors.New("0055")
	ErrNoAssetsFound                            = errors.New("0056")
	ErrNoSegmentsFound                          = errors.New("0057")
	ErrNoPortfoliosFound                        = errors.New("0058")
	ErrNoOrganizationsFound                     = errors.New("0059")
	ErrNoLedgersFound                           = errors.New("0060")
	ErrBalanceUpdateFailed                      = errors.New("0061")
	ErrNoAccountIDsProvided                     = errors.New("0062")
	ErrFailedToRetrieveAccountsByAliases        = errors.New("0063")
	ErrNoAccountsFound                          = errors.New("0064")
	ErrInvalidPathParameter                     = errors.New("0065")
	ErrInvalidAccountType                       = errors.New("0066")
	ErrInvalidMetadataNesting                   = errors.New("0067")
	ErrOperationIDNotFound                      = errors.New("0068")
	ErrNoOperationsFound                        = errors.New("0069")
	ErrTransactionIDNotFound                    = errors.New("0070")
	ErrNoTransactionsFound                      = errors.New("0071")
	ErrInvalidTransactionType                   = errors.New("0072")
	ErrTransactionValueMismatch                 = errors.New("0073")
	ErrForbiddenExternalAccountManipulation     = errors.New("0074")
	ErrAuditRecordNotRetrieved                  = errors.New("0075")
	ErrAuditTreeRecordNotFound                  = errors.New("0076")
	ErrInvalidDateFormat                        = errors.New("0077")
	ErrInvalidFinalDate                         = errors.New("0078")
	ErrDateRangeExceedsLimit                    = errors.New("0079")
	ErrPaginationLimitExceeded                  = errors.New("0080")
	ErrInvalidSortOrder                         = errors.New("0081")
	ErrInvalidQueryParameter                    = errors.New("0082")
	ErrInvalidDateRange                         = errors.New("0083")
	ErrIdempotencyKey                           = errors.New("0084")
	ErrAccountAliasNotFound                     = errors.New("0085")
	ErrLockVersionAccountBalance                = errors.New("0086")
	ErrTransactionIDHasAlreadyParentTransaction = errors.New("0087")
	ErrTransactionIDIsAlreadyARevert            = errors.New("0088")
	ErrTransactionCantRevert                    = errors.New("0089")
	ErrTransactionAmbiguous                     = errors.New("0090")
	ErrBalanceIDNotFound                        = errors.New("0091")
	ErrNoBalancesFound                          = errors.New("0092")
	ErrBalancesCantDeleted                      = errors.New("0093")
)
