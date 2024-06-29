package models

import (
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "time"
)

type Job struct {
    Title           string    `json:"title"`
    Description     string    `json:"description"`
    PostedOn        time.Time `json:"posted_on"`
    TotalApplications int     `json:"total_applications"`
    CompanyName     string    `json:"company_name"`
    PostedBy        string    `json:"posted_by"`
}

func (j *Job) CreateJob() error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    collection := client.Database("recruitment").Collection("jobs")
    _, err := collection.InsertOne(ctx, j)
    return err
}

func GetJob(jobID string) (*Job, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    collection := client.Database("recruitment").Collection("jobs")
    var job Job
    err := collection.FindOne(ctx, bson.M{"_id": jobID}).Decode(&job)
    return &job, err
}

func GetApplicants() ([]User, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    collection := client.Database("recruitment").Collection("users")
    cursor, err := collection.Find(ctx, bson.M{"usertype": "Applicant"})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var applicants []User
    if err = cursor.All(ctx, &applicants); err != nil {
        return nil, err
    }
    return applicants, nil
}

func GetApplicant(applicantID string) (*Profile, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    collection := client.Database("recruitment").Collection("profiles")
    var profile Profile
    err := collection.FindOne(ctx, bson.M{"applicant": applicantID}).Decode(&profile)
    return &profile, err
}

func GetJobs() ([]Job, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    collection := client.Database("recruitment").Collection("jobs")
    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var jobs []Job
    if err = cursor.All(ctx, &jobs); err != nil {
        return nil, err
    }
    return jobs, nil
}

func ApplyToJob(applicantEmail, jobID string) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    collection := client.Database("recruitment").Collection("jobs")
    _, err := collection.UpdateOne(
        ctx,
        bson.M{"_id": jobID},
        bson.M{"$inc": bson.M{"total_applications": 1}},
    )
    return err
}
