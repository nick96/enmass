package main

import (
	"github.com/aryann/difflib"
	"github.com/golang/mock/gomock"
	"github.com/nick96/enmass/mocks"
	"github.com/nick96/enmass/peopleapi"
	"strings"
	"testing"
	"strconv"
	"errors"
)

//go:generate mockgen -destination=mocks/writer_mocks.go -package=mocks io  Writer

func TestGetGroupStat(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	grpMock := mocks.NewMockContactGroup(ctrl)
	writerMock := mocks.NewMockWriter(ctrl)

	var members []peopleapi.Person
	for i := 0; i < 2; i++ {
		member := mocks.NewMockPerson(ctrl)
		member.EXPECT().GetName().Return("Dummy " + strconv.Itoa(i+1))
		member.EXPECT().GetPhoneNumbers().Return([]peopleapi.PhoneNumber{
			peopleapi.PhoneNumber("+44" + strings.Repeat(strconv.Itoa(i+1), 9)),
			peopleapi.PhoneNumber("+61" + strings.Repeat(strconv.Itoa(i+1), 9)),
		})
		member.EXPECT().GetEmails().Return([]peopleapi.Email{
			peopleapi.Email("example" + strconv.Itoa(i+1) + "@example.com"),
			peopleapi.Email("example" + strconv.Itoa(i+1) + "@example.org"),
		})
		members = append(members, member)
	}

	grpMock.EXPECT().GetMembers().Return(members)
	grpMock.EXPECT().GetName().Return("Group Name")
	grpMock.EXPECT().GetGroupID().Return(peopleapi.GroupID("123"))
	grpMock.EXPECT().GetPhoneNumbers().Return([]peopleapi.PhoneNumber{
		"+44123456789",
		"+44123456788",
		"+44123456787",
		"+44123456786",
		"+44123456784",
		"+44123456783",
		"+44123456782",
		"+44123456781",
	})
	grpMock.EXPECT().GetEmails().Return([]peopleapi.Email{
		"example@example.com",
		"example@example.com",
		"example@example.com",
		"example@example.com",
		"example@example.com",
		"example@example.com",
		"example@example.com",
		"example@example.com",
	})
	var written []byte
	writerMock.EXPECT().Write(gomock.Any()).Do(func(p []byte) (int, error) {
		written = append(written, p...)
		return len(p), nil
	}).MinTimes(1)

	err := getGroupStats(grpMock, writerMock)
	if err != nil {
		t.Logf("getGroupStats() returned error")
		t.Fail()
	}

	expected := `Group: Group Name [ID: 123]
2 members 8 phone numbers 8 emails
Dummy 1 (2 phone numbers, 2 emails)
Dummy 2 (2 phone numbers, 2 emails)
`

	if string(written) != expected {
		t.Logf("%s", stringDiff(string(written), expected))
		t.Fail()
	}
}

func TestGetGroupStatSingular(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	grpMock := mocks.NewMockContactGroup(ctrl)
	writerMock := mocks.NewMockWriter(ctrl)

	var members []peopleapi.Person
	for i := 0; i < 2; i++ {
		member := mocks.NewMockPerson(ctrl)
		member.EXPECT().GetName().Return("Dummy " + strconv.Itoa(i+1))
		member.EXPECT().GetPhoneNumbers().Return([]peopleapi.PhoneNumber{
			peopleapi.PhoneNumber("+44" + strings.Repeat(strconv.Itoa(i+1), 9)),
		})
		member.EXPECT().GetEmails().Return([]peopleapi.Email{
			peopleapi.Email("example" + strconv.Itoa(i+1) + "@example.com"),
		})
		members = append(members, member)
	}

	grpMock.EXPECT().GetMembers().Return(members)
	grpMock.EXPECT().GetName().Return("Group Name")
	grpMock.EXPECT().GetGroupID().Return(peopleapi.GroupID("123"))
	grpMock.EXPECT().GetPhoneNumbers().Return([]peopleapi.PhoneNumber{
		"+44123456789",
		"+44123456788",
		"+44123456787",
		"+44123456786",
		"+44123456784",
		"+44123456783",
		"+44123456782",
		"+44123456781",
	})
	grpMock.EXPECT().GetEmails().Return([]peopleapi.Email{
		"example@example.com",
		"example@example.com",
		"example@example.com",
		"example@example.com",
		"example@example.com",
		"example@example.com",
		"example@example.com",
		"example@example.com",
	})
	var written []byte
	writerMock.EXPECT().Write(gomock.Any()).Do(func(p []byte) (int, error) {
		written = append(written, p...)
		return len(p), nil
	}).MinTimes(1)

	err := getGroupStats(grpMock, writerMock)
	if err != nil {
		t.Logf("getGroupStats() returned error")
		t.Fail()
	}

	expected := `Group: Group Name [ID: 123]
2 members 8 phone numbers 8 emails
Dummy 1 (1 phone number, 1 email)
Dummy 2 (1 phone number, 1 email)
`

	if string(written) != expected {
		t.Logf("%s", stringDiff(string(written), expected))
		t.Fail()
	}
}

func TestCheckGroup(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	grpMock := mocks.NewMockContactGroup(ctrl)
	writerMock := mocks.NewMockWriter(ctrl)

	member1 := mocks.NewMockPerson(ctrl)
	member2 := mocks.NewMockPerson(ctrl)
	member3 := mocks.NewMockPerson(ctrl)

	member1.EXPECT().GetEmails().Return([]peopleapi.Email{
		"example@example.com",
	})
	member2.EXPECT().GetEmails().Return([]peopleapi.Email{
		"example@example.com",
	})
	member3.EXPECT().GetEmails().Return([]peopleapi.Email{})

	member1.EXPECT().GetPhoneNumbers().Return([]peopleapi.PhoneNumber{
		"+441232132432",
	})
	member2.EXPECT().GetPhoneNumbers().Return([]peopleapi.PhoneNumber{})
	member3.EXPECT().GetPhoneNumbers().Return([]peopleapi.PhoneNumber{
		"+443243243232",
	})

	member1.EXPECT().GetName().Return("Dummy 1").MinTimes(0)
	member2.EXPECT().GetName().Return("Dummy 2").MinTimes(1)
	member3.EXPECT().GetName().Return("Dummy 3").MinTimes(1)

	members := []peopleapi.Person{member1, member2, member3}

	grpMock.EXPECT().GetMembers().Return(members)

	var written []byte
	writerMock.EXPECT().Write(gomock.Any()).Do(func(p []byte) (int, error) {
		t.Logf("appending: %s", string(p))
		written = append(written, p...)
		t.Logf("written: %s", string(written))
		return len(p), nil
	}).Times(2)

	err := checkGroup(grpMock, writerMock)
	if err != nil {
		t.Logf("checkGroup() return error")
		t.Fail()
	}

	expected := `Dummy 2 has no phone numbers
Dummy 3 has no email
`

	if string(written) != expected {
		t.Logf("\n%s", stringDiff(string(written), expected))
		t.Fail()
	}
}

func TestDoGetActionEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	writerMock0 := mocks.NewMockWriter(ctrl)
	writerMock1 := mocks.NewMockWriter(ctrl)
	grpMock := mocks.NewMockContactGroup(ctrl)

	var written0 []byte
	writerMock0.EXPECT().Write(gomock.Any()).Do(func(p []byte) (int, error) {
		written0 = append(written0, p...)
		return len(p), nil
	})

	var written1 []byte
	writerMock1.EXPECT().Write(gomock.Any()).Do(func(p []byte) (int, error) {
		written1 = append(written1, p...)
		return len(p), nil
	})

	grpMock.EXPECT().GetMembers().Return([]peopleapi.Person{
		mocks.NewMockPerson(ctrl),
		mocks.NewMockPerson(ctrl),
	})
	grpMock.EXPECT().GetEmails().Return([]peopleapi.Email{
		peopleapi.Email("example1@example.com"),
		peopleapi.Email("example2@example.com"),
	})

	err := doGetAction(grpMock, EMAIL, writerMock0, writerMock1)
	if err != nil {
		t.Logf("doGetAction() returned error: %v", err)
		t.Fail()
	}

	expected0 := `Found 2 emails (2 members)
`
	if string(written0) != expected0 {
		t.Logf("\n%s", stringDiff(string(written0), expected0))
		t.Fail()
	}

	expected1 := `example1@example.com; example2@example.com
`
	if string(written1) != expected1 {
		t.Logf("\n%s", stringDiff(string(written1), expected1))
		t.Fail()
	}
}

func TestDoGetActionText(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	writerMock0 := mocks.NewMockWriter(ctrl)
	writerMock1 := mocks.NewMockWriter(ctrl)
	grpMock := mocks.NewMockContactGroup(ctrl)

	var written0 []byte
	writerMock0.EXPECT().Write(gomock.Any()).Do(func(p []byte) (int, error) {
		written0 = append(written0, p...)
		return len(p), nil
	}).MinTimes(0)

	var written1 []byte
	writerMock1.EXPECT().Write(gomock.Any()).Do(func(p []byte) (int, error) {
		written1 = append(written1, p...)
		return len(p), nil
	}).MinTimes(0)

	grpMock.EXPECT().GetMembers().Return([]peopleapi.Person{
		mocks.NewMockPerson(ctrl),
		mocks.NewMockPerson(ctrl),
	})
	grpMock.EXPECT().GetPhoneNumbers().Return([]peopleapi.PhoneNumber{
		peopleapi.PhoneNumber("+32432432434"),
		peopleapi.PhoneNumber("+32432432432"),
	})

	err := doGetAction(grpMock, TEXT, writerMock0, writerMock1)
	if err != nil {
		t.Logf("doGetAction() returned error: %v", err)
		t.Fail()
	}

	expected0 := `Found 2 phone numbers (2 members)
`
	if string(written0) != expected0 {
		t.Logf("\n%s", stringDiff(string(written0), expected0))
		t.Fail()
	}

	expected1 := `+32432432434
+32432432432
`
	if string(written1) != expected1 {
		t.Logf("\n%s", stringDiff(string(written1), expected1))
		t.Fail()
	}
}

func TestSendTextMessageHappy(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockGrp := mocks.NewMockContactGroup(ctrl)
	mockClient := mocks.NewMockClient(ctrl)
	mockWriter := mocks.NewMockWriter(ctrl)

	var written []byte
	mockWriter.EXPECT().Write(gomock.Any()).Do(func (p []byte) (int, error) {
		written = append(written, p...)
		return len(p), nil
	}).MaxTimes(0)

	member := mocks.NewMockPerson(ctrl)
	member.EXPECT().GetPhoneNumbers().Return([]peopleapi.PhoneNumber{
		peopleapi.PhoneNumber("+4343243243232"),
	})
	mockGrp.EXPECT().GetMembers().Return([]peopleapi.Person{member})
	mockClient.EXPECT().SendMessage(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

	err := sendTextMessage(mockGrp, mockClient, "from", "message", mockWriter)
	if err != nil {
		t.Logf("sendTextMessage() returned an error")
		t.Fail()
	}

	if string(written) != "" {
		t.Logf("Unxpected content written: %s", written)
		t.Fail()
	}
}

func TestSendTextMessageUnhappy(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockGrp := mocks.NewMockContactGroup(ctrl)
	mockClient := mocks.NewMockClient(ctrl)
	mockWriter := mocks.NewMockWriter(ctrl)

	var written []byte
	mockWriter.EXPECT().Write(gomock.Any()).Do(func (p []byte) (int, error) {
		written = append(written, p...)
		return len(p), nil
	}).MaxTimes(1)

	member := mocks.NewMockPerson(ctrl)
	member.EXPECT().GetName().Return("name")
	member.EXPECT().GetPhoneNumbers().Return([]peopleapi.PhoneNumber{
		peopleapi.PhoneNumber("to"),
	})
	mockGrp.EXPECT().GetMembers().Return([]peopleapi.Person{member})
	mockClient.EXPECT().SendMessage("from", "to", "message").Return(errors.New("test"))

	err := sendTextMessage(mockGrp, mockClient, "from", "message", mockWriter)
	if err != nil {
		t.Logf("sendTextMessage() returned an error")
		t.Fail()
	}

	expected := "Failed to send message to name on to: test\n"
	if string(written) != expected {
		t.Logf("\n%s", stringDiff(string(written), expected))
		t.Fail()
	}
}

func stringDiff(s1, s2 string) string {
	diffs := difflib.Diff(strings.Split(s1, "\n"), strings.Split(s2, "\n"))

	var stringDiff []string
	for _, diff := range diffs {
		stringDiff = append(stringDiff, diff.String())
	}

	return strings.Join(stringDiff, "\n")
}
