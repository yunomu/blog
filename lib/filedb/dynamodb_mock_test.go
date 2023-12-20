package filedb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type MockDynamoDB struct {
	QueryPagesFn         func(*dynamodb.QueryInput, func(*dynamodb.QueryOutput, bool) bool) error
	TransactWriteItemsFn func(*dynamodb.TransactWriteItemsInput) (*dynamodb.TransactWriteItemsOutput, error)
}

func (m *MockDynamoDB) BatchExecuteStatement(_ *dynamodb.BatchExecuteStatementInput) (*dynamodb.BatchExecuteStatementOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) BatchExecuteStatementWithContext(_ aws.Context, _ *dynamodb.BatchExecuteStatementInput, _ ...request.Option) (*dynamodb.BatchExecuteStatementOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) BatchExecuteStatementRequest(_ *dynamodb.BatchExecuteStatementInput) (*request.Request, *dynamodb.BatchExecuteStatementOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) BatchGetItem(_ *dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) BatchGetItemWithContext(_ aws.Context, _ *dynamodb.BatchGetItemInput, _ ...request.Option) (*dynamodb.BatchGetItemOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) BatchGetItemRequest(_ *dynamodb.BatchGetItemInput) (*request.Request, *dynamodb.BatchGetItemOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) BatchGetItemPages(_ *dynamodb.BatchGetItemInput, _ func(*dynamodb.BatchGetItemOutput, bool) bool) error {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) BatchGetItemPagesWithContext(_ aws.Context, _ *dynamodb.BatchGetItemInput, _ func(*dynamodb.BatchGetItemOutput, bool) bool, _ ...request.Option) error {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) BatchWriteItem(_ *dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) BatchWriteItemWithContext(_ aws.Context, _ *dynamodb.BatchWriteItemInput, _ ...request.Option) (*dynamodb.BatchWriteItemOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) BatchWriteItemRequest(_ *dynamodb.BatchWriteItemInput) (*request.Request, *dynamodb.BatchWriteItemOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) CreateBackup(_ *dynamodb.CreateBackupInput) (*dynamodb.CreateBackupOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) CreateBackupWithContext(_ aws.Context, _ *dynamodb.CreateBackupInput, _ ...request.Option) (*dynamodb.CreateBackupOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) CreateBackupRequest(_ *dynamodb.CreateBackupInput) (*request.Request, *dynamodb.CreateBackupOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) CreateGlobalTable(_ *dynamodb.CreateGlobalTableInput) (*dynamodb.CreateGlobalTableOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) CreateGlobalTableWithContext(_ aws.Context, _ *dynamodb.CreateGlobalTableInput, _ ...request.Option) (*dynamodb.CreateGlobalTableOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) CreateGlobalTableRequest(_ *dynamodb.CreateGlobalTableInput) (*request.Request, *dynamodb.CreateGlobalTableOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) CreateTable(_ *dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) CreateTableWithContext(_ aws.Context, _ *dynamodb.CreateTableInput, _ ...request.Option) (*dynamodb.CreateTableOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) CreateTableRequest(_ *dynamodb.CreateTableInput) (*request.Request, *dynamodb.CreateTableOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DeleteBackup(_ *dynamodb.DeleteBackupInput) (*dynamodb.DeleteBackupOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DeleteBackupWithContext(_ aws.Context, _ *dynamodb.DeleteBackupInput, _ ...request.Option) (*dynamodb.DeleteBackupOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DeleteBackupRequest(_ *dynamodb.DeleteBackupInput) (*request.Request, *dynamodb.DeleteBackupOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DeleteItem(_ *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DeleteItemWithContext(_ aws.Context, _ *dynamodb.DeleteItemInput, _ ...request.Option) (*dynamodb.DeleteItemOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DeleteItemRequest(_ *dynamodb.DeleteItemInput) (*request.Request, *dynamodb.DeleteItemOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DeleteTable(_ *dynamodb.DeleteTableInput) (*dynamodb.DeleteTableOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DeleteTableWithContext(_ aws.Context, _ *dynamodb.DeleteTableInput, _ ...request.Option) (*dynamodb.DeleteTableOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DeleteTableRequest(_ *dynamodb.DeleteTableInput) (*request.Request, *dynamodb.DeleteTableOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeBackup(_ *dynamodb.DescribeBackupInput) (*dynamodb.DescribeBackupOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeBackupWithContext(_ aws.Context, _ *dynamodb.DescribeBackupInput, _ ...request.Option) (*dynamodb.DescribeBackupOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeBackupRequest(_ *dynamodb.DescribeBackupInput) (*request.Request, *dynamodb.DescribeBackupOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeContinuousBackups(_ *dynamodb.DescribeContinuousBackupsInput) (*dynamodb.DescribeContinuousBackupsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeContinuousBackupsWithContext(_ aws.Context, _ *dynamodb.DescribeContinuousBackupsInput, _ ...request.Option) (*dynamodb.DescribeContinuousBackupsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeContinuousBackupsRequest(_ *dynamodb.DescribeContinuousBackupsInput) (*request.Request, *dynamodb.DescribeContinuousBackupsOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeContributorInsights(_ *dynamodb.DescribeContributorInsightsInput) (*dynamodb.DescribeContributorInsightsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeContributorInsightsWithContext(_ aws.Context, _ *dynamodb.DescribeContributorInsightsInput, _ ...request.Option) (*dynamodb.DescribeContributorInsightsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeContributorInsightsRequest(_ *dynamodb.DescribeContributorInsightsInput) (*request.Request, *dynamodb.DescribeContributorInsightsOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeEndpoints(_ *dynamodb.DescribeEndpointsInput) (*dynamodb.DescribeEndpointsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeEndpointsWithContext(_ aws.Context, _ *dynamodb.DescribeEndpointsInput, _ ...request.Option) (*dynamodb.DescribeEndpointsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeEndpointsRequest(_ *dynamodb.DescribeEndpointsInput) (*request.Request, *dynamodb.DescribeEndpointsOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeExport(_ *dynamodb.DescribeExportInput) (*dynamodb.DescribeExportOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeExportWithContext(_ aws.Context, _ *dynamodb.DescribeExportInput, _ ...request.Option) (*dynamodb.DescribeExportOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeExportRequest(_ *dynamodb.DescribeExportInput) (*request.Request, *dynamodb.DescribeExportOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeGlobalTable(_ *dynamodb.DescribeGlobalTableInput) (*dynamodb.DescribeGlobalTableOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeGlobalTableWithContext(_ aws.Context, _ *dynamodb.DescribeGlobalTableInput, _ ...request.Option) (*dynamodb.DescribeGlobalTableOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeGlobalTableRequest(_ *dynamodb.DescribeGlobalTableInput) (*request.Request, *dynamodb.DescribeGlobalTableOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeGlobalTableSettings(_ *dynamodb.DescribeGlobalTableSettingsInput) (*dynamodb.DescribeGlobalTableSettingsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeGlobalTableSettingsWithContext(_ aws.Context, _ *dynamodb.DescribeGlobalTableSettingsInput, _ ...request.Option) (*dynamodb.DescribeGlobalTableSettingsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeGlobalTableSettingsRequest(_ *dynamodb.DescribeGlobalTableSettingsInput) (*request.Request, *dynamodb.DescribeGlobalTableSettingsOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeImport(_ *dynamodb.DescribeImportInput) (*dynamodb.DescribeImportOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeImportWithContext(_ aws.Context, _ *dynamodb.DescribeImportInput, _ ...request.Option) (*dynamodb.DescribeImportOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeImportRequest(_ *dynamodb.DescribeImportInput) (*request.Request, *dynamodb.DescribeImportOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeKinesisStreamingDestination(_ *dynamodb.DescribeKinesisStreamingDestinationInput) (*dynamodb.DescribeKinesisStreamingDestinationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeKinesisStreamingDestinationWithContext(_ aws.Context, _ *dynamodb.DescribeKinesisStreamingDestinationInput, _ ...request.Option) (*dynamodb.DescribeKinesisStreamingDestinationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeKinesisStreamingDestinationRequest(_ *dynamodb.DescribeKinesisStreamingDestinationInput) (*request.Request, *dynamodb.DescribeKinesisStreamingDestinationOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeLimits(_ *dynamodb.DescribeLimitsInput) (*dynamodb.DescribeLimitsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeLimitsWithContext(_ aws.Context, _ *dynamodb.DescribeLimitsInput, _ ...request.Option) (*dynamodb.DescribeLimitsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeLimitsRequest(_ *dynamodb.DescribeLimitsInput) (*request.Request, *dynamodb.DescribeLimitsOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeTable(_ *dynamodb.DescribeTableInput) (*dynamodb.DescribeTableOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeTableWithContext(_ aws.Context, _ *dynamodb.DescribeTableInput, _ ...request.Option) (*dynamodb.DescribeTableOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeTableRequest(_ *dynamodb.DescribeTableInput) (*request.Request, *dynamodb.DescribeTableOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeTableReplicaAutoScaling(_ *dynamodb.DescribeTableReplicaAutoScalingInput) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeTableReplicaAutoScalingWithContext(_ aws.Context, _ *dynamodb.DescribeTableReplicaAutoScalingInput, _ ...request.Option) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeTableReplicaAutoScalingRequest(_ *dynamodb.DescribeTableReplicaAutoScalingInput) (*request.Request, *dynamodb.DescribeTableReplicaAutoScalingOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeTimeToLive(_ *dynamodb.DescribeTimeToLiveInput) (*dynamodb.DescribeTimeToLiveOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeTimeToLiveWithContext(_ aws.Context, _ *dynamodb.DescribeTimeToLiveInput, _ ...request.Option) (*dynamodb.DescribeTimeToLiveOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DescribeTimeToLiveRequest(_ *dynamodb.DescribeTimeToLiveInput) (*request.Request, *dynamodb.DescribeTimeToLiveOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DisableKinesisStreamingDestination(_ *dynamodb.DisableKinesisStreamingDestinationInput) (*dynamodb.DisableKinesisStreamingDestinationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DisableKinesisStreamingDestinationWithContext(_ aws.Context, _ *dynamodb.DisableKinesisStreamingDestinationInput, _ ...request.Option) (*dynamodb.DisableKinesisStreamingDestinationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) DisableKinesisStreamingDestinationRequest(_ *dynamodb.DisableKinesisStreamingDestinationInput) (*request.Request, *dynamodb.DisableKinesisStreamingDestinationOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) EnableKinesisStreamingDestination(_ *dynamodb.EnableKinesisStreamingDestinationInput) (*dynamodb.EnableKinesisStreamingDestinationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) EnableKinesisStreamingDestinationWithContext(_ aws.Context, _ *dynamodb.EnableKinesisStreamingDestinationInput, _ ...request.Option) (*dynamodb.EnableKinesisStreamingDestinationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) EnableKinesisStreamingDestinationRequest(_ *dynamodb.EnableKinesisStreamingDestinationInput) (*request.Request, *dynamodb.EnableKinesisStreamingDestinationOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ExecuteStatement(_ *dynamodb.ExecuteStatementInput) (*dynamodb.ExecuteStatementOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ExecuteStatementWithContext(_ aws.Context, _ *dynamodb.ExecuteStatementInput, _ ...request.Option) (*dynamodb.ExecuteStatementOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ExecuteStatementRequest(_ *dynamodb.ExecuteStatementInput) (*request.Request, *dynamodb.ExecuteStatementOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ExecuteTransaction(_ *dynamodb.ExecuteTransactionInput) (*dynamodb.ExecuteTransactionOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ExecuteTransactionWithContext(_ aws.Context, _ *dynamodb.ExecuteTransactionInput, _ ...request.Option) (*dynamodb.ExecuteTransactionOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ExecuteTransactionRequest(_ *dynamodb.ExecuteTransactionInput) (*request.Request, *dynamodb.ExecuteTransactionOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ExportTableToPointInTime(_ *dynamodb.ExportTableToPointInTimeInput) (*dynamodb.ExportTableToPointInTimeOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ExportTableToPointInTimeWithContext(_ aws.Context, _ *dynamodb.ExportTableToPointInTimeInput, _ ...request.Option) (*dynamodb.ExportTableToPointInTimeOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ExportTableToPointInTimeRequest(_ *dynamodb.ExportTableToPointInTimeInput) (*request.Request, *dynamodb.ExportTableToPointInTimeOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) GetItem(_ *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) GetItemWithContext(_ aws.Context, _ *dynamodb.GetItemInput, _ ...request.Option) (*dynamodb.GetItemOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) GetItemRequest(_ *dynamodb.GetItemInput) (*request.Request, *dynamodb.GetItemOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ImportTable(_ *dynamodb.ImportTableInput) (*dynamodb.ImportTableOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ImportTableWithContext(_ aws.Context, _ *dynamodb.ImportTableInput, _ ...request.Option) (*dynamodb.ImportTableOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ImportTableRequest(_ *dynamodb.ImportTableInput) (*request.Request, *dynamodb.ImportTableOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ListBackups(_ *dynamodb.ListBackupsInput) (*dynamodb.ListBackupsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ListBackupsWithContext(_ aws.Context, _ *dynamodb.ListBackupsInput, _ ...request.Option) (*dynamodb.ListBackupsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ListBackupsRequest(_ *dynamodb.ListBackupsInput) (*request.Request, *dynamodb.ListBackupsOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ListContributorInsights(_ *dynamodb.ListContributorInsightsInput) (*dynamodb.ListContributorInsightsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ListContributorInsightsWithContext(_ aws.Context, _ *dynamodb.ListContributorInsightsInput, _ ...request.Option) (*dynamodb.ListContributorInsightsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ListContributorInsightsRequest(_ *dynamodb.ListContributorInsightsInput) (*request.Request, *dynamodb.ListContributorInsightsOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ListContributorInsightsPages(_ *dynamodb.ListContributorInsightsInput, _ func(*dynamodb.ListContributorInsightsOutput, bool) bool) error {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ListContributorInsightsPagesWithContext(_ aws.Context, _ *dynamodb.ListContributorInsightsInput, _ func(*dynamodb.ListContributorInsightsOutput, bool) bool, _ ...request.Option) error {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ListExports(_ *dynamodb.ListExportsInput) (*dynamodb.ListExportsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ListExportsWithContext(_ aws.Context, _ *dynamodb.ListExportsInput, _ ...request.Option) (*dynamodb.ListExportsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ListExportsRequest(_ *dynamodb.ListExportsInput) (*request.Request, *dynamodb.ListExportsOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ListExportsPages(_ *dynamodb.ListExportsInput, _ func(*dynamodb.ListExportsOutput, bool) bool) error {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ListExportsPagesWithContext(_ aws.Context, _ *dynamodb.ListExportsInput, _ func(*dynamodb.ListExportsOutput, bool) bool, _ ...request.Option) error {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ListGlobalTables(_ *dynamodb.ListGlobalTablesInput) (*dynamodb.ListGlobalTablesOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ListGlobalTablesWithContext(_ aws.Context, _ *dynamodb.ListGlobalTablesInput, _ ...request.Option) (*dynamodb.ListGlobalTablesOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ListGlobalTablesRequest(_ *dynamodb.ListGlobalTablesInput) (*request.Request, *dynamodb.ListGlobalTablesOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ListImports(_ *dynamodb.ListImportsInput) (*dynamodb.ListImportsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ListImportsWithContext(_ aws.Context, _ *dynamodb.ListImportsInput, _ ...request.Option) (*dynamodb.ListImportsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ListImportsRequest(_ *dynamodb.ListImportsInput) (*request.Request, *dynamodb.ListImportsOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ListImportsPages(_ *dynamodb.ListImportsInput, _ func(*dynamodb.ListImportsOutput, bool) bool) error {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ListImportsPagesWithContext(_ aws.Context, _ *dynamodb.ListImportsInput, _ func(*dynamodb.ListImportsOutput, bool) bool, _ ...request.Option) error {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ListTables(_ *dynamodb.ListTablesInput) (*dynamodb.ListTablesOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ListTablesWithContext(_ aws.Context, _ *dynamodb.ListTablesInput, _ ...request.Option) (*dynamodb.ListTablesOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ListTablesRequest(_ *dynamodb.ListTablesInput) (*request.Request, *dynamodb.ListTablesOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ListTablesPages(_ *dynamodb.ListTablesInput, _ func(*dynamodb.ListTablesOutput, bool) bool) error {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ListTablesPagesWithContext(_ aws.Context, _ *dynamodb.ListTablesInput, _ func(*dynamodb.ListTablesOutput, bool) bool, _ ...request.Option) error {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ListTagsOfResource(_ *dynamodb.ListTagsOfResourceInput) (*dynamodb.ListTagsOfResourceOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ListTagsOfResourceWithContext(_ aws.Context, _ *dynamodb.ListTagsOfResourceInput, _ ...request.Option) (*dynamodb.ListTagsOfResourceOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ListTagsOfResourceRequest(_ *dynamodb.ListTagsOfResourceInput) (*request.Request, *dynamodb.ListTagsOfResourceOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) PutItem(_ *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) PutItemWithContext(_ aws.Context, _ *dynamodb.PutItemInput, _ ...request.Option) (*dynamodb.PutItemOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) PutItemRequest(_ *dynamodb.PutItemInput) (*request.Request, *dynamodb.PutItemOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) Query(_ *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) QueryWithContext(_ aws.Context, _ *dynamodb.QueryInput, _ ...request.Option) (*dynamodb.QueryOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) QueryRequest(_ *dynamodb.QueryInput) (*request.Request, *dynamodb.QueryOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) QueryPages(_ *dynamodb.QueryInput, _ func(*dynamodb.QueryOutput, bool) bool) error {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) QueryPagesWithContext(_ aws.Context, in *dynamodb.QueryInput, fn func(*dynamodb.QueryOutput, bool) bool, _ ...request.Option) error {
	if f := m.QueryPagesFn; f != nil {
		return f(in, fn)
	}
	panic("not assigned")
}

func (m *MockDynamoDB) RestoreTableFromBackup(_ *dynamodb.RestoreTableFromBackupInput) (*dynamodb.RestoreTableFromBackupOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) RestoreTableFromBackupWithContext(_ aws.Context, _ *dynamodb.RestoreTableFromBackupInput, _ ...request.Option) (*dynamodb.RestoreTableFromBackupOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) RestoreTableFromBackupRequest(_ *dynamodb.RestoreTableFromBackupInput) (*request.Request, *dynamodb.RestoreTableFromBackupOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) RestoreTableToPointInTime(_ *dynamodb.RestoreTableToPointInTimeInput) (*dynamodb.RestoreTableToPointInTimeOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) RestoreTableToPointInTimeWithContext(_ aws.Context, _ *dynamodb.RestoreTableToPointInTimeInput, _ ...request.Option) (*dynamodb.RestoreTableToPointInTimeOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) RestoreTableToPointInTimeRequest(_ *dynamodb.RestoreTableToPointInTimeInput) (*request.Request, *dynamodb.RestoreTableToPointInTimeOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) Scan(_ *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ScanWithContext(_ aws.Context, _ *dynamodb.ScanInput, _ ...request.Option) (*dynamodb.ScanOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ScanRequest(_ *dynamodb.ScanInput) (*request.Request, *dynamodb.ScanOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ScanPages(_ *dynamodb.ScanInput, _ func(*dynamodb.ScanOutput, bool) bool) error {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) ScanPagesWithContext(_ aws.Context, _ *dynamodb.ScanInput, _ func(*dynamodb.ScanOutput, bool) bool, _ ...request.Option) error {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) TagResource(_ *dynamodb.TagResourceInput) (*dynamodb.TagResourceOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) TagResourceWithContext(_ aws.Context, _ *dynamodb.TagResourceInput, _ ...request.Option) (*dynamodb.TagResourceOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) TagResourceRequest(_ *dynamodb.TagResourceInput) (*request.Request, *dynamodb.TagResourceOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) TransactGetItems(_ *dynamodb.TransactGetItemsInput) (*dynamodb.TransactGetItemsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) TransactGetItemsWithContext(_ aws.Context, _ *dynamodb.TransactGetItemsInput, _ ...request.Option) (*dynamodb.TransactGetItemsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) TransactGetItemsRequest(_ *dynamodb.TransactGetItemsInput) (*request.Request, *dynamodb.TransactGetItemsOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) TransactWriteItems(_ *dynamodb.TransactWriteItemsInput) (*dynamodb.TransactWriteItemsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) TransactWriteItemsWithContext(_ aws.Context, in *dynamodb.TransactWriteItemsInput, _ ...request.Option) (*dynamodb.TransactWriteItemsOutput, error) {
	if f := m.TransactWriteItemsFn; f != nil {
		return f(in)
	}
	panic("not assigned")
}

func (m *MockDynamoDB) TransactWriteItemsRequest(_ *dynamodb.TransactWriteItemsInput) (*request.Request, *dynamodb.TransactWriteItemsOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) UntagResource(_ *dynamodb.UntagResourceInput) (*dynamodb.UntagResourceOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) UntagResourceWithContext(_ aws.Context, _ *dynamodb.UntagResourceInput, _ ...request.Option) (*dynamodb.UntagResourceOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) UntagResourceRequest(_ *dynamodb.UntagResourceInput) (*request.Request, *dynamodb.UntagResourceOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) UpdateContinuousBackups(_ *dynamodb.UpdateContinuousBackupsInput) (*dynamodb.UpdateContinuousBackupsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) UpdateContinuousBackupsWithContext(_ aws.Context, _ *dynamodb.UpdateContinuousBackupsInput, _ ...request.Option) (*dynamodb.UpdateContinuousBackupsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) UpdateContinuousBackupsRequest(_ *dynamodb.UpdateContinuousBackupsInput) (*request.Request, *dynamodb.UpdateContinuousBackupsOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) UpdateContributorInsights(_ *dynamodb.UpdateContributorInsightsInput) (*dynamodb.UpdateContributorInsightsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) UpdateContributorInsightsWithContext(_ aws.Context, _ *dynamodb.UpdateContributorInsightsInput, _ ...request.Option) (*dynamodb.UpdateContributorInsightsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) UpdateContributorInsightsRequest(_ *dynamodb.UpdateContributorInsightsInput) (*request.Request, *dynamodb.UpdateContributorInsightsOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) UpdateGlobalTable(_ *dynamodb.UpdateGlobalTableInput) (*dynamodb.UpdateGlobalTableOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) UpdateGlobalTableWithContext(_ aws.Context, _ *dynamodb.UpdateGlobalTableInput, _ ...request.Option) (*dynamodb.UpdateGlobalTableOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) UpdateGlobalTableRequest(_ *dynamodb.UpdateGlobalTableInput) (*request.Request, *dynamodb.UpdateGlobalTableOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) UpdateGlobalTableSettings(_ *dynamodb.UpdateGlobalTableSettingsInput) (*dynamodb.UpdateGlobalTableSettingsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) UpdateGlobalTableSettingsWithContext(_ aws.Context, _ *dynamodb.UpdateGlobalTableSettingsInput, _ ...request.Option) (*dynamodb.UpdateGlobalTableSettingsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) UpdateGlobalTableSettingsRequest(_ *dynamodb.UpdateGlobalTableSettingsInput) (*request.Request, *dynamodb.UpdateGlobalTableSettingsOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) UpdateItem(_ *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) UpdateItemWithContext(_ aws.Context, _ *dynamodb.UpdateItemInput, _ ...request.Option) (*dynamodb.UpdateItemOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) UpdateItemRequest(_ *dynamodb.UpdateItemInput) (*request.Request, *dynamodb.UpdateItemOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) UpdateTable(_ *dynamodb.UpdateTableInput) (*dynamodb.UpdateTableOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) UpdateTableWithContext(_ aws.Context, _ *dynamodb.UpdateTableInput, _ ...request.Option) (*dynamodb.UpdateTableOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) UpdateTableRequest(_ *dynamodb.UpdateTableInput) (*request.Request, *dynamodb.UpdateTableOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) UpdateTableReplicaAutoScaling(_ *dynamodb.UpdateTableReplicaAutoScalingInput) (*dynamodb.UpdateTableReplicaAutoScalingOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) UpdateTableReplicaAutoScalingWithContext(_ aws.Context, _ *dynamodb.UpdateTableReplicaAutoScalingInput, _ ...request.Option) (*dynamodb.UpdateTableReplicaAutoScalingOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) UpdateTableReplicaAutoScalingRequest(_ *dynamodb.UpdateTableReplicaAutoScalingInput) (*request.Request, *dynamodb.UpdateTableReplicaAutoScalingOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) UpdateTimeToLive(_ *dynamodb.UpdateTimeToLiveInput) (*dynamodb.UpdateTimeToLiveOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) UpdateTimeToLiveWithContext(_ aws.Context, _ *dynamodb.UpdateTimeToLiveInput, _ ...request.Option) (*dynamodb.UpdateTimeToLiveOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) UpdateTimeToLiveRequest(_ *dynamodb.UpdateTimeToLiveInput) (*request.Request, *dynamodb.UpdateTimeToLiveOutput) {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) WaitUntilTableExists(_ *dynamodb.DescribeTableInput) error {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) WaitUntilTableExistsWithContext(_ aws.Context, _ *dynamodb.DescribeTableInput, _ ...request.WaiterOption) error {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) WaitUntilTableNotExists(_ *dynamodb.DescribeTableInput) error {
	panic("not implemented") // TODO: Implement
}

func (m *MockDynamoDB) WaitUntilTableNotExistsWithContext(_ aws.Context, _ *dynamodb.DescribeTableInput, _ ...request.WaiterOption) error {
	panic("not implemented") // TODO: Implement
}
